package user

import (
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

// Delete handles delete user request
func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In User_Delete func")
	u.Dao.DeleteUser(&User{})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted"))
}
