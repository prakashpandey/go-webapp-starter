package user

// Dao interface contains all those methods required for performing database operations.
// All database should implement this interface.
type Dao interface {
	InsertUser(u *User) error
	DeleteUser(u *User) error
}
