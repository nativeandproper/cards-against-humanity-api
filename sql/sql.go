package sql

import (
	"database/sql"
	// Postgres driver
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// NewSQLClient returns SQL instance
func NewSQLClient(sqlAddress string) (*sql.DB, error) {
	db, err := sql.Open("postgres", sqlAddress)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Print("Successfully connected to Postgres")
	return db, nil
}
