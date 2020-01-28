package giphaux

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/backend/sqlite"
	"github.com/redbo/giphaux/shared"
	"github.com/redbo/giphaux/templates"
)

// server holds the server instance's shared state and is the top-level HTTP handler.
type server struct {
	ds          shared.DataStore
	router      *mux.Router
	templates   *template.Template
	gifsDir     string
	tmpDir      string
	uploadLimit int64
}

// userKey serves as a map key for the user's information in the request's Context
var userKey = &struct{}{}

// getUser extracts the user struct associated with this request from the context
func getUser(ctx context.Context) *shared.User {
	if v, ok := ctx.Value(userKey).(*shared.User); ok && v != nil {
		return v
	}
	return nil
}

// templateReload is a hack to enable for more rapid development.
func (s *server) templateReload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.templates, _ = s.templates.ParseGlob("templates/*.tmpl")
		next.ServeHTTP(w, r)
	})
}

// checkAPIKey is a middleware that authorizes a user based on the api_key query parameter.
func (s *server) authorizeAPIUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getUser(r.Context()) == nil {
			s.apiResponse(w, http.StatusUnauthorized, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// checkCookie is a middleware that authorizes a user based on the "session" cookie set in their browser.
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *shared.User
		cookie, err := r.Cookie("session")
		if err == nil {
			user, _ = s.ds.GetUserByCookie(cookie.Value)
		}
		if user == nil {
			apiKey := r.URL.Query().Get("api_key")
			if apiKey != "" {
				user, _ = s.ds.GetUserByKey(apiKey)
			}
		}
		ctx := r.Context()
		if user != nil {
			ctx = context.WithValue(ctx, userKey, user)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// authorizeWebUser is a middleware that makes sure that a user is logged in before they access web pages.
func (s *server) authorizeWebUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getUser(r.Context()) == nil {
			s.error(w, r, http.StatusUnauthorized, "Not Authorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// template executes a template, filling in standard data.
func (s *server) template(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	t := s.templates.Lookup(name)
	err := t.Execute(w, map[string]interface{}{
		// this passes a lot of sensitive user info to all of the templates all the time, which feels like a large surface.
		// something to think about later.
		"User":  getUser(r.Context()),
		"Query": r.URL.Query().Get("q"),
		"Data":  data,
	})
	if err != nil {
		log.Printf("Error executing template: %v", err.Error())
	}
}

// error returns an error page.
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, msg string) {
	w.WriteHeader(code)
	t := s.templates.Lookup("error.tmpl")
	t.Execute(w, map[string]interface{}{
		"User": getUser(r.Context()),
		"Msg":  msg,
		"Data": map[string]interface{}{},
	})
}

// apiResponse responds with serialized JSON.  It automatically appends a Meta object containing the response status.
func (s *server) apiResponse(w http.ResponseWriter, code int, rsp map[string]interface{}) {
	if rsp == nil || len(rsp) == 0 {
		rsp = map[string]interface{}{
			"meta": shared.Meta{Status: code, Msg: http.StatusText(code)},
		}
	}
	js, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

// NewServer creates a new instance of the server.
func NewServer(settings *shared.Configuration) (http.Handler, error) {
	ds, err := sqlite.OpenStore(settings)
	if err != nil {
		return nil, err
	}
	ts, err := templates.LoadTemplates()
	if err != nil {
		return nil, err
	}

	// create the top-level router.
	r := mux.NewRouter()
	s := &server{
		router:      r,
		templates:   ts,
		ds:          ds,
		gifsDir:     settings.GifsDir,
		tmpDir:      settings.TempDir,
		uploadLimit: settings.UploadLimit,
	}

	// Routes for anonymous-accessible web pages.
	r.HandleFunc("/", s.frontPage)
	r.HandleFunc("/login", s.login).Methods("POST")
	r.HandleFunc("/register", s.register).Methods("POST")
	r.HandleFunc("/bye", s.logout)
	r.HandleFunc("/gifs/{id}", s.gifPage)
	r.HandleFunc("/rawgifs/{id:[a-zA-Z0-9]+}.gif", s.rawGif)
	r.HandleFunc("/search", s.search)

	// Routes for API - api_key param must be verified.
	apiAuthed := r.PathPrefix("/v1").Subrouter()
	apiAuthed.HandleFunc("/randomid", nil)
	apiAuthed.HandleFunc("/gifs/search", s.apiSearch)
	apiAuthed.HandleFunc("/gifs/trending", nil)
	apiAuthed.HandleFunc("/gifs/translate", nil)
	apiAuthed.HandleFunc("/gifs/random", nil)
	apiAuthed.HandleFunc("/gifs/{id:[a-zA-Z0-9]+}", s.apiGifID)
	apiAuthed.HandleFunc("/gifs", s.apiGifs)
	apiAuthed.Use(s.authorizeAPIUser)

	// Routes for authenticated web pages - cookie must be verified.
	webAuthed := r.PathPrefix("/user").Subrouter()
	webAuthed.HandleFunc("/", s.userIndex)
	webAuthed.HandleFunc("/upload", s.userUpload).Methods("POST")
	webAuthed.HandleFunc("/delete", s.userDelete).Methods("POST")
	webAuthed.HandleFunc("/favorite", s.userFavorite).Methods("POST")
	webAuthed.HandleFunc("/unfavorite", s.userUnfavorite).Methods("POST")
	webAuthed.HandleFunc("/addcategory", s.userAddCategory).Methods("POST")
	webAuthed.HandleFunc("/removecategory", s.userRemoveCategory).Methods("POST")
	webAuthed.HandleFunc("/updatecategories", s.userUpdateCategories).Methods("POST")
	webAuthed.Use(s.authorizeWebUser)

	r.Use(s.authenticateUser)
	r.Use(s.templateReload) // TODO: TEMPORARY REMOVE THIS
	r.Use(func(n http.Handler) http.Handler {
		return handlers.CombinedLoggingHandler(os.Stdout, n)
	}) // install a middleware that logs requests to stdout in combined apache format

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Run starts the server.
func Run(settings *shared.Configuration) {
	rand.Seed(time.Now().Unix())
	s, err := NewServer(settings)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Listen and serve on 0.0.0.0:8080
	log.Fatal((&http.Server{
		// having i/o timeouts is always a good idea.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         settings.Bind,
		Handler:      s,
	}).ListenAndServe())
}
