package main

import (
	"fmt"
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/auth"
)

// DispatchMethods routes our apis
func DispatchMethods(m map[string]http.HandlerFunc) http.HandlerFunc {
	fmt.Println("In dispatchMethods func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method[%s]\n", r.Method)
		if fn, ok := m[r.Method]; ok {
			fn(w, r)
			return
		}
		MethodNotAllowedHandler(w, r)
	})
}

// MethodNotAllowedHandler handles requests that do not allow the requested method
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In methodNotAllowedHandler func")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}

// UnAuthorizedHandler handles un-authorized requests
func UnAuthorizedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In unAuthorizedHandler func")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Un-Authorized"))
}

// MustAuth authenticates all protected routes
func MustAuth(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	fmt.Println("In mustAuth func")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Validate("elon@spacex.com", "WayToMars") {
			fn(w, r)
			return
		}
		UnAuthorizedHandler(w, r)
	})
}
