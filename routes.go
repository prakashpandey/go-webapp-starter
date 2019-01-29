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
		http.MethodPost: userHandler.Create,
		http.MethodGet:  MustAuth(userHandler.Get),
	}))

	// admin routes
	http.HandleFunc("/admin", DispatchMethods(map[string]http.HandlerFunc{
		http.MethodDelete: MustAuth(adminHandler.DeleteUser),
	}))
}
