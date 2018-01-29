package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// NewSQLClient returns SQL instance
func NewSQLClient(sqlAddress string) (*sql.DB, error) {
	db, err := sql.Open("postgres", sqlAddress)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to Postgres")
	return db, nil
}
