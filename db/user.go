package db

import (
	"log"

	"github.com/prakashpandey/sample-go-webapp/user"
)

// InsertUser in database
func (db DB) InsertUser(u *user.User) error {
	log.Printf("I am creating a user")
	// store user in database here
	return nil
}

// Other func and methods
