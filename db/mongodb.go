// All mongodb related code should be here

package db

import (
	"log"

	"github.com/prakashpandey/sample-go-webapp/user"
)

// DB data structure
type DB struct {
	//
	// make mongodb object part of this struct
	//
	// other fields here
}

// NewDB creates new instance of database
func NewDB() DB {
	return DB{}
}

// InsertUser in database
func (db DB) InsertUser(u *user.User) error {
	// store user in database here
	return nil
}

// DeleteUser in database
func (db DB) DeleteUser(u *user.User) error {
	// store user in database here
	log.Printf("I am deleting user.")
	return nil
}

// Other func and methods
