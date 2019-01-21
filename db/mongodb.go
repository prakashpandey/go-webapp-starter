// All mongodb related code should be here

package db

import (
	"github.com/prakashpandey/sample-go-webapp/user"
)

// Mongodb data structure
type Mongodb struct {
	//
	// make mongodb object part of this struct
	//
	// other fields here
}

// InsertUser in database
func (m Mongodb) InsertUser(u *user.User) error {
	// store user in database here
	return nil
}

// Other func and methods
