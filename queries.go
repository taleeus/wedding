package main

import (
	"database/sql"
	"errors"
)

const findGuest = /* sql */ `
SELECT
    rowid AS id,
    name,
    surname,
    answer,
    answered_at,
    created_at
FROM guest
WHERE rowid = ?
`

func FindGuest(db *sql.DB, id int) (Guest, bool, error) {
	var guest Guest

	row := db.QueryRow(findGuest, id)
	if err := row.Scan(
		&guest.ID,
		&guest.Name,
		&guest.Surname,
		&guest.Answer,
		&guest.AnsweredAt,
		&guest.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Guest{}, false, nil
		}

		return Guest{}, false, err
	}

	return guest, true, nil
}
