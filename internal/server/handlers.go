package server

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/taleeus/wedding/internal/db"
	"github.com/taleeus/wedding/web/pages"
)

func SaveRSVP(w http.ResponseWriter, r *http.Request, dbconn *sql.DB) {
	rsvp := db.RSVP{
		Phone:   sanitizePhone(r.FormValue("phone")),
		Name:    sanitize(r.FormValue("name")),
		Surname: sanitize(r.FormValue("surname")),
		Guests:  nullable(sanitize(r.FormValue("guests"))),
		Food:    nullable(sanitize(r.FormValue("food"))),
		Notes:   nullable(sanitize(r.FormValue("notes"))),
	}

	success := true
	if err := db.InsertRSVP(r.Context(), dbconn, rsvp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		success = false
	}

	if err := pages.Feedback(success).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sanitize(str string) string {
	return strings.TrimSpace(str)
}

func sanitizePhone(phone string) string {
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "+39", "")

	return phone
}

func nullable(str string) sql.Null[string] {
	return sql.Null[string]{
		V:     str,
		Valid: str != "",
	}
}
