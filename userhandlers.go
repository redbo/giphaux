package main

import "net/http"

func (s *server) userIndex(w http.ResponseWriter, r *http.Request) {
	t := s.templates.Lookup("user.tmpl")
	t.Execute(w, nil)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
