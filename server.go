package giphaux

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/backend/sqlite"
	"github.com/redbo/giphaux/shared"
	"github.com/redbo/giphaux/templates"
	"go.uber.org/zap"
)

// server holds the server instance's shared state and is the top-level HTTP handler.
type server struct {
	ds          shared.DataStore
	router      *mux.Router
	templates   *template.Template
	gifsDir     string
	tmpDir      string
	uploadLimit int64
	queryLimit  int
	logger      *zap.Logger
}

var ( // these serve as map keys for items in a request's Context
	userKey   = &struct{}{}
	loggerKey = &struct{}{}
)

// getUser extracts the user struct associated with this request from the context
func getUser(ctx context.Context) *shared.User {
	if v, ok := ctx.Value(userKey).(*shared.User); ok && v != nil {
		return v
	}
	return nil
}

// log returns the request's logger, which is basically scoped with the request ID
func (s *server) log(r *http.Request) *zap.Logger {
	if v, ok := r.Context().Value(loggerKey).(*zap.Logger); ok && v != nil {
		return v
	}
	return s.logger // if the request doesn't have a logger, just return the server logger.
}

// checkAPIKey is a middleware that authorizes a user based on the api_key query parameter.
func (s *server) authorizeAPIUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getUser(r.Context()) == nil {
			s.log(r).Info("API user not authorized")
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
			s.log(r).Info("User not authorized")
			s.error(w, r, http.StatusUnauthorized, "Not Authorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// logMiddleware adds a logger with a unique request ID to the request context, and also logs accesses.
func (s *server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := s.logger.With(
			zap.String("RequestID", fmt.Sprintf("%x", rand.Int63())),
		)
		ctx := context.WithValue(r.Context(), loggerKey, logger)
		m := httpsnoop.CaptureMetrics(next, w, r.WithContext(ctx))
		logger.Info("ACCESS",
			zap.String("method", r.Method),
			zap.String("remote_address", r.RemoteAddr),
			zap.String("referer", r.Referer()),
			zap.String("agent", r.UserAgent()),
			zap.String("path", r.URL.String()),
			zap.Time("start", start),
			zap.Int("response_code", m.Code),
			zap.Int64("response_size", m.Written),
			zap.Duration("duration", m.Duration),
		)
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
		s.logger.Error("Error executing template: %v", zap.Error(err))
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
		s.logger.Error("Error marshalling json", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

// NewServer creates a new instance of the server.
func NewServer(settings *shared.Configuration, logger *zap.Logger) (http.Handler, error) {
	ds, err := sqlite.OpenStore(settings, logger)
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
		queryLimit:  settings.MaxQueryLimit,
		logger:      logger,
	}

	// Routes for anonymous-accessible web pages.
	r.HandleFunc("/", s.frontPage)
	r.HandleFunc("/login", s.login).Methods("POST")
	r.HandleFunc("/register", s.register).Methods("POST")
	r.HandleFunc("/bye", s.logout)
	r.HandleFunc("/gifs/{id}", s.gifPage)
	r.HandleFunc("/gif/{id:[a-zA-Z0-9]+}.gif", s.rawGif)
	r.HandleFunc("/still/{id:[a-zA-Z0-9]+}.gif", s.stillGif)
	r.HandleFunc("/search", s.search)

	// Routes for API - api_key param must be verified.
	apiAuthed := r.PathPrefix("/v1").Subrouter()
	apiAuthed.HandleFunc("/randomid", s.apiRandomID)
	apiAuthed.HandleFunc("/gifs/search", s.apiSearch)
	apiAuthed.HandleFunc("/gifs/trending", s.apiTrending)
	apiAuthed.HandleFunc("/gifs/translate", s.apiTranslate)
	apiAuthed.HandleFunc("/gifs/random", s.apiRandomSearch)
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

	r.Use(s.logMiddleware)
	r.Use(s.authenticateUser)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.templates, _ = s.templates.ParseGlob("templates/*.tmpl")
			next.ServeHTTP(w, r)
		})
	}) // TODO: TEMPORARY DEV STUFF, REMOVE THIS

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Run starts the server.
func Run(settings *shared.Configuration) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := NewServer(settings, logger)
	if err != nil {
		logger.Fatal("Error creating server", zap.Error(err))
	}

	log.Fatal((&http.Server{
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		Addr:         settings.Bind,
		Handler:      s,
	}).ListenAndServe())
}
