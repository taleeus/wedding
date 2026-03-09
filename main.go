package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dburl = "$TURSO_DATABASE_URL?authToken=$TURSO_AUTH_TOKEN"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env not found; using environment")
	}

	db, err := sql.Open("libsql", os.ExpandEnv(dburl))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	ctx := context.Background()
	if err := migrate(ctx, db); err != nil {
		panic(err)
	}
}
