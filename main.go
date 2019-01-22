package main

import (
	"log"
	"net/http"
)

func main() {
	routes()
	url := ":8284"
	log.Printf("Started webserver at %s", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
