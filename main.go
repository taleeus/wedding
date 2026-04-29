package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/taleeus/wedding/internal/db"
	"github.com/taleeus/wedding/internal/server"
)

var dburl = "libsql://$TURSO_DATABASE_URL?authToken=$TURSO_AUTH_TOKEN"

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})))

	err := godotenv.Load()
	if err != nil {
		slog.Info(".env not found; using environment")
	}

	dbconn, err := db.New(os.ExpandEnv(dburl))
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()

	if err := db.Migrate(dbconn); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("🚀 starting server", "port", port)
	if err := http.ListenAndServe(":"+port, server.New(dbconn)); err != nil {
		log.Fatal(err)
	}
}
