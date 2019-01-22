package main

import (
	"log"
	"net/http"
)

// Server contains shared resources
type Server struct {
	//db db.Database
	// other shared resources here
}

func main() {
	s := &Server{
		//db: mongodb.Mongodb{
		//	// Initialize the real mongodb object here
		//},
	}
	s.routes()
	url := ":8284"
	log.Printf("Started webserver at %s", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
