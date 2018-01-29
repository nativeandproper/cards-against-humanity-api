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
			token VARCHAR(32) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			verified_at TIMESTAMP,
			expires_at TIMESTAMP NOT NULL
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
