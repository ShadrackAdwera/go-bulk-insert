package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeData(ctx context.Context, payload *[]CovidData, opts ...asynq.Option) error
}

type DataTaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(clientOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(clientOpts)
	return &DataTaskDistributor{
		client,
	}
}
