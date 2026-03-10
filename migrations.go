package main

import (
	"database/sql"
	"fmt"
)

var migrations = []string{
	debugTable,
}

func migrate(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %w", err)
	}
	defer tx.Rollback()

	for i, stmt := range migrations {
		if _, err := tx.Exec(stmt); err != nil {
			return fmt.Errorf("error executing statement %d: %w", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing migrations: %w", err)
	}

	return nil
}

const debugTable = /* sql */ `
CREATE TABLE IF NOT EXISTS debug (
    message TEXT NOT NULL
);
`
