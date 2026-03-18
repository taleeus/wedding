package main

import (
	"database/sql"
	"fmt"
	"time"
)

const schema = /* sql */ `
CREATE TABLE IF NOT EXISTS debug (
    message TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS guest (
    name        TEXT        NOT NULL,
    surname     TEXT        NOT NULL,
    answer      TEXT                    CHECK (answer IN ('YES', 'NO', 'MAYBE')),
    answered_at DATETIME,
    created_at  DATETIME    NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS guest_name_idx
ON guest(LOWER(name), LOWER(surname));
`

type Answer string

const (
	AnswerYes   Answer = "YES"
	AnswerNo    Answer = "NO"
	AnswerMaybe Answer = "MAYBE"
)

type Guest struct {
	ID            int
	Name, Surname string
	Answer        sql.Null[Answer]
	AnsweredAt    sql.Null[time.Time]
	CreatedAt     time.Time
}

func initDB(db *sql.DB) error {
	tx, err := db.Begin()
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
