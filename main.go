package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

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

func debugHandler(tmpl *template.Template, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guestIDStr := r.URL.Query().Get("g")
		if guestIDStr == "" {
			http.NotFound(w, r)
			return
		}

		guestID, err := strconv.Atoi(guestIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		guest, ok, err := FindGuest(db, guestID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			http.NotFound(w, r)
			return
		}

		data := TmplData{Title: "Titolo", Msg: guest.Name + " " + guest.Surname}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
