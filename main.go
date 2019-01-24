package main

import (
	"log"
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/db"
	"github.com/prakashpandey/sample-go-webapp/user"
)

// DB will be almost same for every handler
var DB = db.NewDB()

// Add all handlers here
var userHandler = user.Handler{
	Dao: DB,
}

func main() {
	routes()
	url := ":8284"
	log.Printf("Starting webserver at %s", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
