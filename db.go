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
