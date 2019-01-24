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
	http.HandleFunc("/user", dispatchMethods(map[string]http.HandlerFunc{
		"POST":   userHandler.Create,
		"DELETE": userHandler.Delete,
	},
	))
}
