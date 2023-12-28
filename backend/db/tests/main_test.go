package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/m3phist/gobank/backend/db/sqlc"
	"github.com/m3phist/gobank/backend/utils"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load env config", err)
	}

	conn, err := sql.Open(config.DB_driver, config.DB_source)
	if err != nil {
		log.Fatal("Could not connect to database", err)

	}

	testQuery = db.New(conn)

	os.Exit(m.Run())
}
