package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-boilerplate/config"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/otel"
)

type Task struct {
	taskClient    *asynq.Client
	taskInspector *asynq.Inspector
}

func NewTask(config *config.Config) *Task {
	redisAddr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	taskClient := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, Password: config.Redis.Password})
	taskInspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: redisAddr, Password: config.Redis.Password})

	return &Task{
		taskClient,
		taskInspector,
	}
}

func (t *Task) CreateTask(ctx context.Context, taskName string, payload interface{}, queue string, retry int) error {
	ctx, span := otel.Start(ctx, "worker:task:CreateTask", map[string]interface{}{
		"taskName": taskName,
		"payload":  payload,
		"queue":    queue,
		"retry":    retry,
	})
	defer span.End()

	payloadB, error := json.Marshal(payload)
	if error != nil {
		return error
	}

	info, err := t.taskClient.Enqueue(asynq.NewTask(taskName, payloadB), asynq.Queue(queue), asynq.MaxRetry(retry))
	if err != nil {
		return err
	}

	logger.Info(ctx).Msgf("Task created: id=%s name=%s payload=%s queue=%s", info.ID, info.Type, info.Payload, info.Queue)
	return nil
}

func (t *Task) ArchiveScheduledTasks(queue string) error {
	tasks, err := t.taskInspector.ArchiveAllScheduledTasks(queue)
	if err != nil {
		return err
	}

	logger.Info(context.Background()).Msgf("Archived %d scheduled tasks", tasks)
	return nil
}

func (t *Task) DeleteAllScheduledTasks(queue string) error {
	tasks, err := t.taskInspector.DeleteAllScheduledTasks(queue)
	if err != nil {
		return err
	}

	logger.Info(context.Background()).Msgf("Deleted %d scheduled tasks", tasks)
	return nil
}
