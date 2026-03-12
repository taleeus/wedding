package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dburl = "$TURSO_DATABASE_URL?authToken=$TURSO_AUTH_TOKEN"

//go:embed views
var views embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env not found; using environment")
	}

	db, err := sql.Open("libsql", os.ExpandEnv(dburl))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initDB(db); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	tmpl, err := template.ParseFS(views, "views/debug.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", debugHandler(tmpl, db))

	log.Printf("running server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type TmplData struct {
	Title string
	Msg   string
}

func debugHandler(tmpl *template.Template, db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var message string
		row := db.QueryRow("SELECT message FROM debug")
		if err := row.Scan(&message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := TmplData{Title: "Titolo", Msg: message}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
