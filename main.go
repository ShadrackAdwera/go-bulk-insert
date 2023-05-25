package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ShadrackAdwera/go-bulk-insert/api"
	db "github.com/ShadrackAdwera/go-bulk-insert/db/sqlc"
	"github.com/ShadrackAdwera/go-bulk-insert/worker"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	zerolog "github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	url := os.Getenv("DEV_DB")
	// check if valid URL
	conn, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf(err.Error())
	}

	redisAddr := os.Getenv("DEV_DB")
	// check if valid redis address
	clientOpts := asynq.RedisClientOpt{
		Addr: redisAddr,
	}

	distro := worker.NewTaskDistributor(clientOpts)

	store := db.NewStore(conn)

	srv := api.NewServer(distro, store)

	serverAddr := os.Getenv("SERVER_ADDRESS")

	srv.StartServer(serverAddr)

	go startTaskProcessor(clientOpts, store)
}

func startTaskProcessor(opts asynq.RedisClientOpt, store db.TxStore) {
	processor := worker.NewTaskServer(opts, store)

	err := processor.Start()

	if err != nil {
		zerolog.Fatal().Err(err).Msg("error starting the redis task processor")
		return
	}
	zerolog.Info().Msg("redis task processor started . . . ")
}
