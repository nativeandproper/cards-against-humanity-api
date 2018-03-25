package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00003, Down00003)
}

func Up00003(tx *sql.Tx) error {
	// create user_verification_tokens table
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS user_verification_tokens (
			id SERIAL PRIMARY KEY, 
			user_id INTEGER NOT NULL REFERENCES users(id), 
			token VARCHAR(44) UNIQUE NOT NULL,
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
			verified_at timestamp without time zone,
			expires_at timestamp without time zone NOT NULL
		   )
		`)
	if err != nil {
		return err
	}
	return nil
}

func Down00003(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user_verification_tokens")
	if err != nil {
		return err
	}

	return nil
}
