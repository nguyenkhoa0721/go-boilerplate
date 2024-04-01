package redis

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	GoRedis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go-boilerplate/config"
)

type RedisClient struct {
	client        *GoRedis.Client
	taskClient    *asynq.Client
	taskInspector *asynq.Inspector
}

func NewRedisClient(config *config.Config) (*RedisClient, error) {
	redisAddr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	client := GoRedis.NewClient(&GoRedis.Options{
		Addr:     redisAddr,
		Password: config.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	logrus.Info("Redis connected")

	taskClient := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, Password: config.Redis.Password})
	taskInspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: redisAddr, Password: config.Redis.Password})

	return &RedisClient{
		client,
		taskClient,
		taskInspector,
	}, nil
}

func (r *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *RedisClient) SetValue(ctx context.Context, key string, val string) error {
	err := r.client.Set(ctx, key, val, 0).Err()

	return err
}

func (r *RedisClient) IncValue(ctx context.Context, key string) error {
	err := r.client.Incr(ctx, key).Err()

	return err
}
