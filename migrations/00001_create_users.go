package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUsers, downUsers)
}

func upUsers(tx *sql.Tx) error {
	query := `CREATE TABLE users (
		"id" bigserial not null primary key,
		"email" varchar not null unique,
		"encrypted_password" varchar not null
	);`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func downUsers(tx *sql.Tx) error {
	query := `DROP TABLE users;`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
