// Provides structs and functions for interacting with the thread model
package data

import "time"

// Struct representing a thread
type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}
