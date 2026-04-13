package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/go-libsql"
)

func New(dsn string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
