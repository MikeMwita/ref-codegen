// frameworks/database/db.go

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// NewDB creates a new database connection
func NewDB() (*sql.DB, error) {
	// get the database configuration from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// create the database connection string
	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	// open the database connection
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	// ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// return the database connection
	return db, nil
}
