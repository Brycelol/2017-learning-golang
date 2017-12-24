// Provides structs and functions for interacting with the user model
package data

import "time"

// Struct representing a user
type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Struct representing a user session
type Session struct {
	Id        string
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
