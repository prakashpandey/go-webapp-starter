package main

import (
	"fmt"
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/auth"
)

func dispatchMethods(m map[string]http.HandlerFunc) http.HandlerFunc {
	fmt.Println("In dispatchMethods func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method[%s]\n", r.Method)
		if fn, ok := m[r.Method]; ok {
			mustAuth(fn)(w, r)
			return
		}
		methodNotAllowedHandler()(w, r)
	})
}

func methodNotAllowedHandler() http.HandlerFunc {
	fmt.Println("In methodNotAllowedHandler func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}

func unAuthorizedHandler() http.HandlerFunc {
	fmt.Println("In unAuthorizedHandler func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Un-Authorized"))
	})
}

// authenticate all protected routes
func mustAuth(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	fmt.Println("In mustAuth func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Validate("elon@spacex.com", "WayToMars") {
			fn(w, r)
			return
		}
		unAuthorizedHandler()
	})
}
