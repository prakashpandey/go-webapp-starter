package main

import (
	"log"
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/auth"
	"github.com/prakashpandey/sample-go-webapp/user"
)

func (s *Server) routes() {
	// define all routes here
	http.HandleFunc("/", s.mustAuth(s.rootHandler))
}

// authanticate all protected routes
func (s *Server) mustAuth(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Validate(user.Username, user.Password) {
			fn(w, r)
			return
		}
		w.Write([]byte("Un-Authorised user"))
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Hi from root handler")); err != nil {
		log.Printf("error: %s", err.Error())
	}
}
