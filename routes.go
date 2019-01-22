package main

import (
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/auth"
	"github.com/prakashpandey/sample-go-webapp/db"
	"github.com/prakashpandey/sample-go-webapp/index"
	"github.com/prakashpandey/sample-go-webapp/user"
)

// DB will be almost same for every handler
var DB = db.NewDB()

// Add all handlers here
var userHandler = user.Handler{
	Dao: DB,
}

func routes() {
	// define all routes here
	http.HandleFunc("/", index.HelloHandler)
	http.HandleFunc("/user", mustAuth(user.CreateUserHandler))
	http.HandleFunc("/user/delete", mustAuth(userHandler.DeleteUserHandler))
}

// authenticate all protected routes
func mustAuth(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Validate("elon@spacex.com", "WayToMars") {
			fn(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Un-Authorized user"))
		return
	})
}
