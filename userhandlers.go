package giphaux

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/redbo/giphaux/shared"
	"go.uber.org/zap"
)

// These handlers are meant for web browser access and are only accessible if you're logged in.

// userIndex is the front landing page for the user.
func (s *server) userIndex(w http.ResponseWriter, r *http.Request) {
	user := getUser(r.Context())
	fp, err := s.ds.UserFrontpage(user.Username)
	if err != nil {
		s.log(r).Error("Error getting user homepage", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error finding user")
		return
	}
	s.template(w, r, "user.tmpl", fp)
}

// search is the HTML page that displays search results for a given query.
func (s *server) userUploads(w http.ResponseWriter, r *http.Request) {
	user := getUser(r.Context())
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}
	limit := 12
	gifs, totalresults, err := s.ds.UserUploads(user.Username, limit, offset)
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	data := map[string]interface{}{ // build a datastructure to pass to the template
		"Gifs":         gifs,
		"Gifcount":     len(gifs),
		"TotalResults": totalresults,
		"PrevOffset":   offset - limit,
		"NextOffset":   offset + len(gifs),
		"Offset":       offset,
		"Limit":        limit,
		"Start":        offset + 1,
	}
	s.template(w, r, "search.tmpl", data)
}

// userAddCategory is the handler for a user adding a category.
func (s *server) userAddCategory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	category, err := shared.NormalizeTag(r.FormValue("category"))
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "Invalid category name")
		return
	}
	user := getUser(r.Context())
	if err := s.ds.AddCategory(user.Username, category); err != nil {
		s.log(r).Error("Error creating category", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Unable to create category")
		return
	}
	http.Redirect(w, r, "/user/#Categories", http.StatusSeeOther)
}

// userRemoveCategory is the handler for a user removing a category.
func (s *server) userRemoveCategory(w http.ResponseWriter, r *http.Request) {
	user := getUser(r.Context())
	r.ParseForm()
	category, err := shared.NormalizeTag(r.FormValue("category"))
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "Invalid category name")
		return
	}
	if err := s.ds.RemoveCategory(user.Username, category); err != nil {
		s.log(r).Error("Error removing category", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error removing category.")
		return
	}
	http.Redirect(w, r, "/user/#Categories", http.StatusSeeOther)
}

func (s *server) userUpdateCategories(w http.ResponseWriter, r *http.Request) {
	user := getUser(r.Context())
	gif := r.FormValue("gifid")
	r.ParseForm()
	cats := []string{}
	for key := range r.Form {
		if strings.HasPrefix(key, "cat-") {
			if nk, err := shared.NormalizeTag(key[4:]); err == nil {
				cats = append(cats, nk)
			}
		}
	}
	if err := s.ds.UpdateCategories(user.Username, gif, cats); err != nil {
		s.log(r).Error("Error updating categories", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error updating categories")
		return
	}
	http.Redirect(w, r, "/gifs/"+gif, http.StatusSeeOther)
}

// userFavorite is the handler for a user favoriting a GIF.
func (s *server) userFavorite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	gif := r.FormValue("gifid")
	user := getUser(r.Context())
	cats := []string{}
	for key := range r.Form {
		if strings.HasPrefix(key, "cat-") {
			if nk, err := shared.NormalizeTag(key[4:]); err == nil {
				cats = append(cats, nk)
			}
		}
	}
	if err := s.ds.AddFavorite(user.Username, gif, cats); err != nil {
		s.log(r).Error("Error favoriting gif", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error favoriting GIF")
		return
	}
	http.Redirect(w, r, "/gifs/"+gif, http.StatusSeeOther)
}

// userUnfavorite is the handler for a user unfavoriting a GIF.
func (s *server) userUnfavorite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	gif := r.FormValue("gifid")
	user := getUser(r.Context())
	if err := s.ds.RemoveFavorite(user.Username, gif); err != nil {
		s.log(r).Error("Error unfavoriting gif", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error unfavoriting GIF")
		return
	}
	http.Redirect(w, r, "/gifs/"+gif, http.StatusSeeOther)
}

// userUpload is the handler for a user uploading a new GIF.
func (s *server) userUpload(w http.ResponseWriter, r *http.Request) {
	var rating string
	r.ParseMultipartForm(s.uploadLimit)
	user := getUser(r.Context())
	if user == nil {
		s.log(r).Error("No user?")
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		s.log(r).Error("Error getting gif upload", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error saving the file")
		return
	}
	defer file.Close()
	filedata, err := ioutil.ReadAll(file)
	if err != nil {
		s.log(r).Error("Error reading upload", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error reading uploaded gif")
		return
	}
	width, height, frames, size, err := shared.GIFInfo(filedata)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, "Error parsing gif")
		return
	}
	caption := r.FormValue("caption")
	tags := make([]string, 0)
	for _, tag := range strings.Split(r.FormValue("tags"), ",") {
		if nt, err := shared.NormalizeTag(tag); err == nil {
			tags = append(tags, nt)
		}
	}
	sourceURL := ""
	if rating, err = shared.NormalizeRating(r.FormValue("rating")); err != nil {
		rating = "g"
	}
	cats := []string{}
	for key := range r.Form {
		if strings.HasPrefix(key, "cat-") {
			if nk, err := shared.NormalizeTag(key[4:]); err == nil {
				cats = append(cats, nk)
			}
		}
	}

	gif, err := s.ds.AddGIF(user.Username, caption, tags, cats, sourceURL, rating, width, height, size, frames, filedata)
	if err != nil {
		s.log(r).Error("Error saving gif to database", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error persisting gif to database")
		return
	}

	http.Redirect(w, r, "/gifs/"+gif.ID, http.StatusSeeOther)
}

// userDelete is the handler for a user deleting a GIF.
func (s *server) userDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	gifid, err := shared.NormalizeGIFID(r.FormValue("gifid"))
	if err != nil {
		s.error(w, r, http.StatusNotFound, "Unable to find that image.")
		return
	}
	user := getUser(r.Context())
	if err := s.ds.RemoveGIF(user.Username, gifid); err != nil {
		s.log(r).Error("Error deleting GIF", zap.Error(err))
		s.error(w, r, http.StatusInternalServerError, "Error deleting GIF")
		return
	}
	http.Redirect(w, r, "/user/", http.StatusSeeOther)
}
