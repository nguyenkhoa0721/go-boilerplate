package worker

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go-boilerplate/config"
	"go-boilerplate/pkg/constant"
)

type Worker struct {
	mux *asynq.ServeMux
	srv *asynq.Server
}

func NewWorker(config *config.Config) *Worker {
	redisAddr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Password: config.Redis.Password,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				constant.CRITICAL_QUEUE: 6,
				constant.DEFAULT_QUEUE:  3,
				constant.LOW_QUEUE:      1,
			},
		},
	)

	return &Worker{
		mux: asynq.NewServeMux(),
		srv: srv,
	}
}

func (w *Worker) RegisterHandler(pattern string, handler func(context.Context, *asynq.Task) error) {
	w.mux.HandleFunc(pattern, handler)
}

func (w *Worker) StartWorker() error {
	if err := w.srv.Run(w.mux); err != nil {
		return err
	}

	return nil
}
