package giphaux

import (
	"net/http"
	"strconv"

	"github.com/redbo/giphaux/shared"
)

// These handlers are for API access and primarily return json objects.

func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var err error
	q := r.URL.Query().Get("q")
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil || limit < 0 || limit > s.queryLimit {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil || offset < 0 {
		s.apiResponse(w, http.StatusBadRequest, nil)
		return
	}
	gifs, totalResults, err := s.ds.Search(q, limit, offset)
	if err != nil {
		s.apiResponse(w, http.StatusInternalServerError, nil)
		return
	}
	rsp := map[string]interface{}{
		"data": gifs,
		"pagination": shared.Pagination{
			TotalCount: totalResults,
			Count:      len(gifs),
			Offset:     offset,
		},
	}
	s.apiResponse(w, http.StatusOK, rsp)
}

func (s *server) apiGifID(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiGifs(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiRandomID(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiTrending(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiTranslate(w http.ResponseWriter, r *http.Request) {
}

func (s *server) apiRandomSearch(w http.ResponseWriter, r *http.Request) {
}
