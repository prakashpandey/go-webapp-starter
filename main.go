package main

import (
	"log"
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/db"
)

// Server contains shared resources
type Server struct {
	db db.Database
	// other shared resources here
}

func main() {
	s := &Server{
		db: db.Mongodb{
		// Initialize the real mongodb object here
		},
	}
	s.routes()
	url := ":8855"
	log.Printf("Started webserver at %s", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
