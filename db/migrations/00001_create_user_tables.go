package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {

	// create users table
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			first_name VARCHAR(20) NOT NULL, 
			last_name VARCHAR(20) NOT NULL,
			email VARCHAR(40) UNIQUE NOT NULL, 
			password bytea NOT NULL,
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			deleted_at timestamp without time zone 
			)
		  `)
	if err != nil {
		return err
	}

	// create account types enum
	_, err = tx.Exec(`CREATE TYPE account_type AS ENUM ('basic')`)
	if err != nil {
		return err
	}

	// create account types table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS account_types (
			id SERIAL PRIMARY KEY,
			type account_type DEFAULT 'basic',
			request_limit INTEGER NOT NULL,
			api_key_limit INTEGER NOT NULL,
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			updated_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC')
	       )
		`)
	if err != nil {
		return err
	}

	// create user account type history table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS user_account_type_history (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			account_type_id INTEGER NOT NULL REFERENCES account_types(id),
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			expires_at timestamp without time zone NOT NULL
		   )
		`)
	if err != nil {
		return err
	}

	// create api keys table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS api_keys (
			id SERIAL PRIMARY KEY,
			api_key VARCHAR(32) UNIQUE NOT NULL,
			created_at timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
			deleted_at timestamp without time zone 
	       )
		`)
	if err != nil {
		return err
	}

	// create user api keys table
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS user_api_keys (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL references users(id),
		api_key_id INTEGER NOT NULL REFERENCES api_keys(id)
	   )
	`)
	if err != nil {
		return err
	}

	return nil
}

func Down00001(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user_account_type_history")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS user_api_keys")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS api_keys")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS account_types")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TYPE account_type")
	if err != nil {
		return err
	}
	return nil
}
