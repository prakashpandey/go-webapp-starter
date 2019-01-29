package db

import (
	"log"

	"github.com/prakashpandey/sample-go-webapp/user"
)

// DeleteUser in database
func (db DB) DeleteUser(u *user.User) error {
	// store user in database here
	log.Printf("I am deleting user from [db layer]")
	return nil
}
