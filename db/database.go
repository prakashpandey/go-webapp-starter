package db

import (
	"github.com/prakashpandey/sample-go-webapp/user"
)

// Database interface contains all those methods required for performing database operations.
// All database should implement this interface.
type Database interface {
	InsertUser(u *user.User) error
}
