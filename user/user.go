package user

import (
	"log"
	"net/http"
)

// User data structure
type User struct {
	name string
	// Other User fields here
}

// CreateUserHandler handle user request
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Hi from create user handler. This is a secured endpoint")); err != nil {
		log.Printf("error: %s", err.Error())
	}
}

type UserHandler struct {
	Dao Dao
}

// DeleteUserHandler handle user delete request
func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	u.Dao.DeleteUser(&User{})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted"))
}
