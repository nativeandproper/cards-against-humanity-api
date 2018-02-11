package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {

	// create sets table
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS sets (
		id SERIAL PRIMARY KEY, 
		name VARCHAR(50),
		created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC')
	   )
	`)
	if err != nil {
		return err
	}

	// create black cards table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS black_cards (
			id SERIAL PRIMARY KEY, 
			text text NOT NULL, 
			pick INTEGER NOT NULL,
			set_id  INTEGER NOT NULL REFERENCES sets(id), 
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			deleted_at timestamp without time zone
		   )
		`)
	if err != nil {
		return err
	}

	// create white cards table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS white_cards (
		id SERIAL PRIMARY KEY, 
		text text NOT NULL, 
		set_id  INTEGER NOT NULL REFERENCES sets(id), 
		created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		deleted_at timestamp without time zone 
	   )
	`)
	if err != nil {
		return err
	}

	return nil
}

func Down00002(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS white_cards")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS black_cards")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS sets")
	if err != nil {
		return err
	}
	return nil
}
