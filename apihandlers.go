package giphaux

import (
	"net/http"

	"github.com/redbo/giphaux/shared"
)

// These handlers are for API access and primarily return json objects.

func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	limit := 12
	offset := 0
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
	w.Write([]byte("GIF ID PAGE"))
}

func (s *server) apiGifs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GIFS PAGE"))
}
