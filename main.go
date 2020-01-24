package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var userlist = [][3]string{
	[3]string{"test", "test", "testtest123"},
}

type server struct {
	db        *sql.DB
	router    *mux.Router
	templates *template.Template
}

type userNameKey struct{}

func (s *server) checkAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedUser := ""
		apiKey := r.URL.Query().Get("api_key")
		if apiKey != "" {
			for _, user := range userlist {
				if user[2] == apiKey {
					authedUser = user[0]
				}
			}
		}
		if authedUser == "" {
			http.Error(w, "NOT AUTHORIZED", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userNameKey{}, authedUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) checkCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedUser := ""
		cookie, err := r.Cookie("session")
		if err == nil {
			for _, user := range userlist {
				if user[2] == cookie.Value {
					authedUser = user[0]
				}
			}
		}
		if authedUser == "" {
			http.Error(w, "NOT AUTHORIZED", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userNameKey{}, authedUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewServer creates a new instance of the server.
func NewServer() (*server, error) {
	ts, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	// create the top-level router.
	r := mux.NewRouter()
	s := &server{
		router:    r,
		templates: ts,
	}

	// Routes for anonymous-accessible web pages.
	r.HandleFunc("/", s.frontPage)
	r.HandleFunc("/fronts", s.fronts)
	r.HandleFunc("/login", s.login).Methods("POST")
	r.HandleFunc("/loginFailed", s.loginFail)
	r.HandleFunc("/register", s.register).Methods("POST")

	// Routes for API - api_key param must be verified.
	apiAuthed := r.PathPrefix("/v1").Subrouter()
	apiAuthed.HandleFunc("/gifs/search", s.apiSearch)
	apiAuthed.HandleFunc("/stickers/search", nil)
	apiAuthed.HandleFunc("/gifs/trending", nil)
	apiAuthed.HandleFunc("/stickers/trending", nil)
	apiAuthed.HandleFunc("/gifs/translate", nil)
	apiAuthed.HandleFunc("/stickers/translate", nil)
	apiAuthed.HandleFunc("/gifs/random", nil)
	apiAuthed.HandleFunc("/stickers/random", nil)
	apiAuthed.HandleFunc("/randomid", nil)
	apiAuthed.HandleFunc("/gifs/{id:[0-9]+}", s.apiGifID)
	apiAuthed.HandleFunc("/gifs", s.apiGifs)
	apiAuthed.Use(s.checkAPIKey)

	// Routes for authenticated web pages - cookie must be verified.
	webAuthed := r.PathPrefix("/user").Subrouter()
	webAuthed.HandleFunc("/", s.userIndex)
	webAuthed.Use(s.checkCookie)

	return s, nil
}

func main() {
	s, err := NewServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Listen and serve on 0.0.0.0:8080
	log.Fatal((&http.Server{
		// having i/o timeouts is always a good idea.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         "0.0.0.0:8080",
		Handler:      s,
	}).ListenAndServe())
}
