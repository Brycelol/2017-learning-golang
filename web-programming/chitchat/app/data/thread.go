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

// We have attached a method to a pointer of a thread struct
// This give the function access to the thread. In this instance
// the thread struct is known as the 'receiver'
//
// This function will count the number of replies the thread has received
func (thread *Thread) NumReplies() (int, error) {

	// Note the prepared statement here - we inject the thread id
	rows, err := Db.Query("SELECT COUNT(*) FROM posts WHERE thread_id=$1", thread.Id)

	if err != nil {
		return 0, err
	}

	var count int

	for rows.Next() {
		// Scan row (should only be one) and assign it to the count variable
		err := rows.Scan(&count)

		if err != nil {
			return 0, err
		}
	}

	// Close off db connection
	rows.Close()

	// Return the count, and a nil error value as all is good
	return count, nil
}

// Function to return all threads from the database
// This corresponds to a function as it is not attached to any types; pointers or otherwise
// Returns a slice of Threads and/or an error
func Threads() ([]Thread, error) {

	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")

	// Return the error straight away to avoid nesting errors
	// Keeps things readable
	if err != nil {
		return nil, err
	}

	// Declare the slice of threads to be returned
	var threads []Thread

	// Iterate over returned rows
	for rows.Next() {

		// Create a new thread instance to hold the row state
		th := Thread{}

		// Scan copies the columns in the current row into the values pointed
		// at by dest. The number of values in dest must be the same as the
		// number of columns in Rows.
		err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt)

		// Return if an error occurs while scanning the row
		if err != nil {
			return nil, err
		}

		// We retrieved a thread, so append it in
		threads = append(threads, th)
	}

	// Close the resource
	rows.Close()

	// Return our slice of threads, and null as no error
	return threads, nil
}
