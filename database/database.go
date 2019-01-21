package database

// Database interface contains all those methods required for performing database operations.
// All database should implement this interface.
type Database interface {
	InsertUser() error
}
