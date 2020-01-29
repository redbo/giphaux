package giphaux

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/shared"
	"go.uber.org/zap"
)

// These handlers are for API access and primarily return json objects.

func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var rating string
	var err error
	q := r.URL.Query().Get("q")
	if rating, err = shared.NormalizeRating(r.URL.Query().Get("rating")); err != nil {
		rating = "g" // default search to a "g" rating
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit < 0 || limit > s.queryLimit {
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
	s.apiResponse(w, http.StatusOK, map[string]interface{}{
		"data": gifs,
		"pagination": shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

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
	s.apiResponse(w, http.StatusOK, map[string]interface{}{
		"data": gif,
	})
}

func (s *server) apiGifs(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var err error
	gifids := []string{}
	for _, gid := range strings.Split(mux.Vars(r)["ids"], ",") {
		gifid, err := shared.NormalizeGIFID(gid)
		if err == nil {
			gifids = append(gifids, gifid)
		}
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit < 0 || limit > s.queryLimit {
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
	s.apiResponse(w, http.StatusOK, map[string]interface{}{
		"data": gifs,
		"pagination": shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

func (s *server) apiRandomID(w http.ResponseWriter, r *http.Request) {
	s.apiResponse(w, http.StatusOK, map[string]interface{}{
		"data": shared.RandomID{
			ID: s.ds.RandomID(),
		},
	})
}

func (s *server) apiTrending(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var rating string
	var err error
	if rating, err = shared.NormalizeRating(r.URL.Query().Get("rating")); err != nil {
		rating = "g" // default search to a "g" rating
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit < 0 || limit > s.queryLimit {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil || offset < 0 {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	gifs, totalResults, err := s.ds.Trending(limit, offset, rating)
	if err != nil {
		s.log(r).Error("Error getting trending gifs", zap.Error(err))
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	s.apiResponse(w, http.StatusOK, map[string]interface{}{
		"data": gifs,
		"pagination": shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	})
}

func (s *server) apiTranslate(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiRandomSearch(w http.ResponseWriter, r *http.Request) {
}
