package server

import (
	"log/slog"
	"net/http"
)

func SaveGuestResponse(w http.ResponseWriter, req *http.Request) {
	slog.InfoContext(req.Context(), "RSVP", "form", req.Form)
}
