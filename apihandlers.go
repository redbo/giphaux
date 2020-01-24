package main

import (
	"encoding/json"
	"net/http"
)

func apiResponse(w http.ResponseWriter, code int, rsp map[interface{}]interface{}) {
	if rsp == nil {
		rsp = map[interface{}]interface{}{
			"meta": Meta{Status: code, Msg: http.StatusText(code)},
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

func (s *server) apiSearch(w http.ResponseWriter, r *http.Request) {
	gifs := []GIF{}
	rsp := map[interface{}]interface{}{
		"data": gifs,
		"pagination": Pagination{
			TotalCount: len(gifs),
			Count:      len(gifs),
			Offset:     0,
		},
	}
	apiResponse(w, http.StatusOK, rsp)
}

func (s *server) apiGifID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GIF ID PAGE"))
}

func (s *server) apiGifs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GIFS PAGE"))
}
