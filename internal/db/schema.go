package db

import (
	"database/sql"
	"fmt"
)

const schema = /* sql */ `
CREATE TABLE IF NOT EXISTS rsvp (
    phone       TEXT        NOT NULL    PRIMARY KEY,
    name        TEXT        NOT NULL,
    surname     TEXT        NOT NULL,
    guests      TEXT,
    food        TEXT,
    notes       TEXT,
    created_at  DATETIME    NOT NULL    DEFAULT CURRENT_TIMESTAMP
);
`

func Migrate(dbconn *sql.DB) error {
	tx, err := dbconn.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(schema); err != nil {
		return fmt.Errorf("error executing schema statement: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing schema tx: %w", err)
	}

	return nil
}
