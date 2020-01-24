package main

import (
	"fmt"
	"net/http"
	"time"
)

func (s *server) frontPage(w http.ResponseWriter, r *http.Request) {
	t := s.templates.Lookup("index.tmpl")
	t.Execute(w, nil)
}

func (s *server) fronts(w http.ResponseWriter, r *http.Request) {
}

func (s *server) loginFail(w http.ResponseWriter, r *http.Request) {
	t := s.templates.Lookup("error.tmpl")
	t.Execute(w, nil)
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	for _, user := range userlist {
		if user[0] == username && user[1] == password {
			http.SetCookie(w, &http.Cookie{
				Name:    "session",
				Value:   user[2],
				Expires: time.Now().Add(time.Hour),
			})
		}
		fmt.Println(w.Header())
		http.Redirect(w, r, "/user/", 301)
		return
	}
	http.Redirect(w, r, "/failedLogin", 301)
}

func (s *server) register(w http.ResponseWriter, r *http.Request) {
}
