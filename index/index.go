package index

import (
	"log"
	"net/http"
)

// HelloHandler hello world handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Hi from hello handler. This is not secured endpoint")); err != nil {
		log.Printf("error: %s", err.Error())
	}
}
