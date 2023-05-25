package worker

import (
	"context"

	db "github.com/ShadrackAdwera/go-bulk-insert/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskProcessor interface {
	TaskProcessData(ctx context.Context, task *asynq.Task) error
	Start() error
}

type DataTaskProcessor struct {
	server *asynq.Server
	store  db.TxStore
}

func NewTaskServer(opts asynq.RedisClientOpt, store db.TxStore) TaskProcessor {
	server := asynq.NewServer(opts, asynq.Config{
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Err(err).Str("task_type", task.Type()).Bytes("payload", task.Payload()).Msg("error processing task . . ")
		}),
	})
	return &DataTaskProcessor{
		server, store,
	}
}

func (processor *DataTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskInsertData, processor.TaskProcessData)

	return processor.server.Start(mux)
}
