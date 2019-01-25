package main

import (
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/index"
)

func routes() {
	// define all routes here

	// default
	http.HandleFunc("/", index.HelloHandler)

	// user routes
	http.HandleFunc("/user", DispatchMethods(map[string]http.HandlerFunc{
		"POST":   userHandler.Create,
		"DELETE": MustAuth(userHandler.Delete),
	}))
}
