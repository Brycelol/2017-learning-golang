// Provides database access
package data

import (
	"database/sql"
	"log"
)

// Global variable which has an pointer to an sql DB object. Represents a pool
// of database connections
var Db *sql.DB

// Init function - Called at application startup to initalize the database connection pool
func init() {

	// Need to declare an error here as we are assigning an existing Db global object
	var err error

	// Initialise the database variable object
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=false")

	if err != nil {
		log.Fatal(err)
	}
}
