package sql

import (
	"database/sql"
	"github.com/rs/zerolog"
)

type DatabaseClient struct {
	sqlClient *sql.DB
	logger    zerolog.Logger
}

// NewDatabaseClient creates a new instance of the databse client
func NewDatabaseClient(sqlClient *sql.DB, logger zerolog.Logger) *DatabaseClient {
	return &DatabaseClient{
		sqlClient,
		logger,
	}
}
