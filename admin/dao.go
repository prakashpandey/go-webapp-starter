package admin

import (
	"github.com/prakashpandey/sample-go-webapp/user"
)

// Dao interface contains all those methods required for performing admin related database operations.
// All database should implement this interface.
type Dao interface {
	DeleteUser(u *user.User) error
}
