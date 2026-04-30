package server

import (
	"database/sql"
	"net/http"

	"github.com/a-h/templ"
	"github.com/taleeus/wedding/static"
	"github.com/taleeus/wedding/web/pages"
)

func New(dbconn *sql.DB, copy pages.HomeCopy) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", templ.Handler(pages.Home(copy)))
	mux.HandleFunc("POST /", recoverable(extend(SaveRSVP(copy.Feedback), dbconn)))

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static.Assets)))

	return mux
}
