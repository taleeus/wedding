package db

import (
	"context"
	"database/sql"
)

type RSVP struct {
	Name    string
	Surname string
	Phone   string
	Guests  sql.Null[string]
	Food    sql.Null[string]
	Notes   sql.Null[string]
}

const insertRSVP = /* sql */ `
INSERT INTO rsvp
	(phone, name, surname, guests, food, notes) VALUES
	(?, ?, ?, ?, ?, ?)
ON CONFLICT (phone) DO UPDATE SET
	name = excluded.name,
	surname = excluded.surname,
	guests = excluded.guests,
	food = excluded.food,
	notes = excluded.notes
`

func InsertRSVP(ctx context.Context, conn *sql.DB, rsvp RSVP) error {
	_, err := conn.ExecContext(ctx, insertRSVP,
		rsvp.Phone,
		rsvp.Name,
		rsvp.Surname,
		rsvp.Guests,
		rsvp.Food,
		rsvp.Notes,
	)

	return err
}
