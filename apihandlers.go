package giphaux

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/shared"
	"go.uber.org/zap"
)

// These handlers are for API access and primarily return json objects.

// apiSearch returns a list of gifs for the given search term and constraints.
func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var rating string
	var err error
	q := r.URL.Query().Get("q")
	if rating, err = shared.NormalizeRating(r.URL.Query().Get("rating")); err != nil {
		rating = "g" // default search to a "g" rating
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit <= 0 || limit > s.queryLimit {
		limit = s.queryLimit
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil || offset < 0 {
		offset = 0
	}
	gifs, totalResults, err := s.ds.Search(q, limit, offset, rating)
	if err != nil {
		s.log(r).Error("Error searching gifs", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{
		Data: gifs,
		Pagination: &shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

// apiGifID returns a gif given its id.
func (s *server) apiGifID(w http.ResponseWriter, r *http.Request) {
	gifid, err := shared.NormalizeGIFID(mux.Vars(r)["id"])
	if err != nil {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	gif, err := s.ds.GIFByID(gifid)
	if err != nil || gif == nil {
		s.apiResponse(w, http.StatusNotFound, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{Data: gif})
}

// apiGifs returns a list of gifs given their comma-separated IDs.
func (s *server) apiGifs(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var err error
	gifids := []string{}
	for _, gid := range strings.Split(r.URL.Query().Get("ids"), ",") {
		gifid, err := shared.NormalizeGIFID(gid)
		if err == nil {
			gifids = append(gifids, gifid)
		}
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit <= 0 || limit > s.queryLimit {
		limit = s.queryLimit
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil || offset < 0 {
		offset = 0
	}
	gifs, totalResults, err := s.ds.GIFsByID(gifids, limit, offset)
	if err != nil {
		s.log(r).Error("Error finding gifs", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{
		Data: gifs,
		Pagination: &shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

// apiRandomID returns a random ID.
func (s *server) apiRandomID(w http.ResponseWriter, r *http.Request) {
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{
		Data: &shared.RandomID{ID: s.ds.RandomID()},
	})
}

// apiTrending returns a list of trending GIFs.
func (s *server) apiTrending(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var rating string
	var err error
	if rating, err = shared.NormalizeRating(r.URL.Query().Get("rating")); err != nil {
		rating = "g" // default search to a "g" rating
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit < 0 || limit > s.queryLimit {
		limit = s.queryLimit
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil || offset < 0 {
		offset = 0
	}
	gifs, totalResults, err := s.ds.Trending(limit, offset, rating)
	if err != nil {
		s.log(r).Error("Error getting trending gifs", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{
		Data: gifs,
		Pagination: &shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

// apiTranslate searches for a random gif with the given search terms.  weirdness currently has no effect.
func (s *server) apiTranslate(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("s")
	weirdness, err := strconv.Atoi(r.URL.Query().Get("weirdness"))
	if err != nil || weirdness < 0 || weirdness > 10 { // if I implement it, it's definitely going to 11.
		weirdness = 0
	}
	gif, err := s.ds.RandomSearch(q, weirdness)
	if err != nil {
		s.log(r).Error("Error getting random gif", zap.Error(err))
		s.apiResponse(w, http.StatusNotFound, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{Data: gif})
}

// apiRandomSearch searches for a random gif given a tag and rating.
func (s *server) apiRandomSearch(w http.ResponseWriter, r *http.Request) {
	tag, err := shared.NormalizeTag(r.URL.Query().Get("tag"))
	if err != nil {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	rating, err := shared.NormalizeRating(r.URL.Query().Get("tag"))
	if err != nil {
		rating = "g"
	}
	gif, err := s.ds.RandomByTag(tag, rating)
	if err != nil {
		s.log(r).Error("Error getting random gif by tag", zap.Error(err))
		s.apiResponse(w, http.StatusNotFound, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, &shared.APIResponse{Data: gif})
}

// apiUploadGif is the endpoint for uploading gifs to the system.
func (s *server) apiUploadGif(w http.ResponseWriter, r *http.Request) {
	rating := "g" // I don't know how this gets populated on upload API
	caption := "" // ditto
	cats := []string{}
	var filedata []byte
	var sourceURL string
	var file io.ReadCloser
	var err error

	r.ParseMultipartForm(s.uploadLimit)
	user := getUser(r.Context())
	if user == nil {
		s.log(r).Error("No user?")
		return
	}

	if sourceURL = r.FormValue("source_image_url"); sourceURL != "" {
		resp, err := http.Get(sourceURL)
		if err != nil {
			s.apiResponse(w, http.StatusBadRequest, nil)
			return
		}
		file = resp.Body
	} else if file, _, err = r.FormFile("file"); err != nil {
		s.log(r).Error("Error getting gif upload", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	} else {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	defer file.Close()
	if filedata, err = ioutil.ReadAll(file); err != nil {
		s.log(r).Error("Error reading upload", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}

	width, height, size, frames, err := shared.GIFInfo(filedata)
	if err != nil {
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	tags := make([]string, 0)
	for _, tag := range strings.Split(r.FormValue("tags"), ",") {
		if nt, err := shared.NormalizeTag(tag); err == nil {
			tags = append(tags, nt)
		}
	}

	if _, err = s.ds.AddGIF(user.Username, caption, tags, cats, sourceURL,
		rating, width, height, size, frames, filedata); err != nil {
		s.log(r).Error("Error saving gif to database", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}

	s.apiResponse(w, http.StatusOK, nil)
}
