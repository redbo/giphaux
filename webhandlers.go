package giphaux

import (
	"image/gif"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/redbo/giphaux/shared"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// These handlers are meant for web browser access, and are accessbile whether you are logged in or not.

// frontPage is the front page of the website, aka "/"
func (s *server) frontPage(w http.ResponseWriter, r *http.Request) {
	s.template(w, r, "index.tmpl", nil)
}

// login handles the form submssions for a user that is logging in.
func (s *server) login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username, err := shared.NormalizeUsername(r.FormValue("username"))
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Invalid username")
		return
	}
	password := r.FormValue("password")
	if password == "" {
		s.error(w, r, http.StatusUnauthorized, "Invalid password.")
		return
	}
	user, err := s.ds.GetUser(username)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Unable to authenticate user.")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Unable to authenticate user.")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    user.Cookie,
		Expires:  time.Now().Add(time.Hour),
		SameSite: http.SameSiteLaxMode, // enable some browser CSRF protections
	})
	http.Redirect(w, r, "/user/", http.StatusSeeOther)
}

// register handles the form submission for new account creation.
func (s *server) register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// TODO: input validation on username and password
	username, err := shared.NormalizeUsername(r.FormValue("username"))
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Invalid username")
		return
	}
	password := r.FormValue("password")
	if password == "" {
		s.error(w, r, http.StatusUnauthorized, "Invalid password.")
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Unable to create user.")
		return
	}
	user, err := s.ds.NewUser(username, string(bytes))
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, "Unable to create user.")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    user.Cookie,
		Expires:  time.Now().Add(time.Hour),
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/user/", http.StatusSeeOther) // send the user to their personal home page
}

// logout logs a user out of their account.
func (s *server) logout(w http.ResponseWriter, r *http.Request) {
	// overwrite session cookie with an empty one that expires immediately
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// gifPage is the HTML page that displays a single gif and information about it.
func (s *server) gifPage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	gif, err := s.ds.GIFByID(id)
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	var userinfo *shared.UserGIFInfo
	if user := getUser(r.Context()); user != nil {
		userinfo, _ = s.ds.UserGIFInfo(user.Username, id)
	}
	data := map[string]interface{}{ // build a datastructure to pass to the template
		"ID":       mux.Vars(r)["id"],
		"Tags":     gif.Tags,
		"Title":    gif.Caption,
		"UserInfo": userinfo,
	}
	s.template(w, r, "gif.tmpl", data)
}

// rawGif serves the actual gif image with an image/gif content-type.
func (s *server) rawGif(w http.ResponseWriter, r *http.Request) {
	gifid, err := shared.NormalizeGIFID(mux.Vars(r)["id"])
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	// DANGER DANGER DANGER make sure the gifid is normalized before touching the filesystem.
	fp, err := os.Open(filepath.Join(s.gifsDir, gifid+".gif"))
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	defer fp.Close()
	w.Header().Set("Content-Type", "image/gif")
	io.Copy(w, fp)
}

// stillGif serves the first frame of the gif with an image/gif content-type.
func (s *server) stillGif(w http.ResponseWriter, r *http.Request) {
	gifid, err := shared.NormalizeGIFID(mux.Vars(r)["id"])
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	// DANGER DANGER DANGER make sure the gifid is normalized before touching the filesystem.
	fp, err := os.Open(filepath.Join(s.gifsDir, gifid+".gif"))
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	defer fp.Close()

	img, err := gif.DecodeAll(fp)
	if err != nil || len(img.Image) < 1 {
		s.error(w, r, http.StatusInternalServerError, "Error parsing gif")
		return
	}
	w.Header().Set("Content-Type", "image/gif")
	gif.Encode(w, img.Image[0], nil)
}

// search is the HTML page that displays search results for a given query.
func (s *server) search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}
	limit := 12
	gifs, totalresults, err := s.ds.Search(q, limit, offset, "g")
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	data := map[string]interface{}{ // build a datastructure to pass to the template
		"Gifs":         gifs,
		"TotalResults": totalresults,
		"Offset":       offset,
		"Limit":        limit,
		"Query":        q,
	}
	s.template(w, r, "search.tmpl", data)
}
