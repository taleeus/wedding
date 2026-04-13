package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/taleeus/wedding/internal/db"
	"github.com/taleeus/wedding/internal/server"
)

var dburl = "libsql://$TURSO_DATABASE_URL?authToken=$TURSO_AUTH_TOKEN"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env not found; using environment")
	}

	conn, err := db.New(os.ExpandEnv(dburl))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 listening on port %s", port)
	if err := http.ListenAndServe(":"+port, server.New()); err != nil {
		log.Fatal(err)
	}
}
