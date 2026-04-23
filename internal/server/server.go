package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/taleeus/wedding/static"
	"github.com/taleeus/wedding/web/pages"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", templ.Handler(pages.Home()))
	mux.HandleFunc("POST /", SaveGuestResponse)

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static.Assets)))

	return mux
}
