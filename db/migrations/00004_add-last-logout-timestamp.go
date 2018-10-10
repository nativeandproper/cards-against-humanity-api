package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`ALTER TABLE users ADD COLUMN logged_out_at TIMESTAMP`)
	if err != nil {
		return err
	}
	return nil
}

func Down00004(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`ALTER TABLE users DROP COLUMN logged_out_at`)
	if err != nil {
		return err
	}
	return nil
}
