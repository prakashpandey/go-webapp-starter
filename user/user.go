package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User data structure
type User struct {
	// Name of user
	Name string `json:"name"`
	// Other User fields here
}

// Handler sample user handler
type Handler struct {
	Dao Dao
}

// Create handles create user request
func (u *Handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In User_Create func")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Hi from create user handler. This is not a secured endpoint")); err != nil {
		log.Printf("error: %s", err.Error())
	}
}

// Get handles sample get user request
func (u *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In User_Get func")
	user := User{
		Name: "Prakash Pandey",
	}
	var resp []byte
	var err error
	if resp, err = json.Marshal(user); err != nil {
		log.Printf("error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
