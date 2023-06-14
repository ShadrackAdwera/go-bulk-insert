package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQuery TxStore

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), "postgresql://postgres:password@localhost:5431/test_go_bulk_insert?sslmode=disable")

	if err != nil {
		log.Fatalf("Failed to initialize the database %v", err)
	}

	testQuery = NewStore(pool)

	os.Exit(m.Run())
}
