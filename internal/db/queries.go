package db

import (
	"database/sql"
	"errors"
)

const findGuestByFullName = /* sql */ `
SELECT
    rowid AS id,
    name,
    surname,
    answer,
    answered_at,
    created_at
FROM guest
WHERE
    LOWER(name) = ? AND
    LOWER(surname) = ?
`

func FindGuestByFullName(db *sql.DB, name, surname string) (Guest, bool, error) {
	var guest Guest

	row := db.QueryRow(findGuestByFullName, name, surname)
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
