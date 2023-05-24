package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery TxStore

func TestMain(m *testing.M) {
	testDb, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5431/test_go_bulk_insert?sslmode=disable")

	if err != nil {
		log.Fatalf("Failed to initialize the database %v", err)
	}

	testQuery = NewStore(testDb)

	os.Exit(m.Run())
}
