package auth

import (
	"github.com/prakashpandey/sample-go-webapp/user"
)

// THIS is the user which requested the operations
type THIS user.User

// Role of different user
type Role struct {
	CREATE []user.User
	READ   []user.User
	UPDATE []user.User
	DELETE []user.User
}
