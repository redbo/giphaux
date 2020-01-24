package giphaux

import (
	"net/http"

	"github.com/redbo/giphaux/shared"
)

// These handlers are for API access and primarily return json objects.

func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	gifs := []shared.GIF{}
	rsp := map[interface{}]interface{}{
		"data": gifs,
		"pagination": shared.Pagination{
			TotalCount: len(gifs),
			Count:      len(gifs),
			Offset:     0,
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
