package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

// Attempts to connect to the database
//
// Parameters:
//   - connectionString: connection string
//
// Returns:
//   - Database Connection
//   - An error if there is an issue connecting to the databse
func Connect(connectionString string) (*sql.DB, error) {
	if db == nil {
		db, err = sql.Open("mysql", connectionString)
	}

	return db, err
}

// Returns a list of valid origins
//
// Parameters:
//   - None
//
// Returns:
//   - A list of valid origins
func ValidOrigins() []string {
	return []string{"https://staging.benpthom.com", "https://benpthom.com"}
}

// Returns a list of valid component names
//
// Parameters:
//   - None
//
// Returns:
//   - A list of valid component names
func ValidComponents() []string {
	return []string{"About", "Experience", "Skills", "Projects"}
}
