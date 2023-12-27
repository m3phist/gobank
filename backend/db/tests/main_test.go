package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/m3phist/gobank/backend/db/sqlc"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgres://root:9uN9H9FMyhLtgkLi@localhost:5432/gobank_db?sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect to database", err)

	}

	testQuery = db.New(conn)

	os.Exit(m.Run())
}
