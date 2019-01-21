package main

import (
	"net/http"
)

func (s *Server) routes() {
	// define all routes here

	http.HandleFunc("/", s.rootHandler)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi from root handler"))
	w.WriteHeader(http.StatusOK)
	return
}
