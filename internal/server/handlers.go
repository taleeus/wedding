package server

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strings"

	"github.com/taleeus/wedding/internal/db"
	"github.com/taleeus/wedding/web/pages"
)

func SaveRSVP(copy pages.FeedbackCopy) func(w http.ResponseWriter, r *http.Request, dbconn *sql.DB) {
	return func(w http.ResponseWriter, r *http.Request, dbconn *sql.DB) {
		rsvp := db.RSVP{
			Phone:   sanitizePhone(r.FormValue("phone")),
			Name:    sanitize(r.FormValue("name")),
			Surname: sanitize(r.FormValue("surname")),
			Guests:  nullable(sanitize(r.FormValue("guests"))),
			Food:    nullable(sanitize(r.FormValue("food"))),
			Notes:   nullable(sanitize(r.FormValue("notes"))),
		}

		if rsvp.Phone == "" || rsvp.Name == "" || rsvp.Surname == "" {
			slog.WarnContext(r.Context(), "Invalid RSVP",
				"rsvp", rsvp,
			)

			http.Error(w, "mind your own business, useless bot", http.StatusBadRequest)
			return
		}

		success := true
		slog.InfoContext(r.Context(), "Saving response",
			"rsvp", rsvp,
		)

		if err := db.InsertRSVP(r.Context(), dbconn, rsvp); err != nil {
			slog.ErrorContext(r.Context(), "RSVP upserting failed",
				"err", err.Error(),
			)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			success = false
		}

		if err := pages.Feedback(copy, success).Render(r.Context(), w); err != nil {
			slog.ErrorContext(r.Context(), "Feedback rendering failed",
				"err", err.Error(),
			)

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
