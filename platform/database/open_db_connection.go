package database

import (
	"Hrms/app/queries"
	"os"

	"github.com/jmoiron/sqlx"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.EmployeeQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *sqlx.DB
		err error
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		db, err = PostgreSQLConnection()
	}

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		EmployeeQueries: &queries.EmployeeQueries{DB: db},
	}, nil
}
