package uuid

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/pkg/constant"
	"go-boilerplate/pkg/database/redis"
	"go-boilerplate/pkg/logger"
	"strconv"
	"time"
)

const (
	COUNTER_KEY = "uuid_counter"
	COUNTER_MAX = 4294967295
)

type Uuid struct {
	min         uint64
	max         uint64
	nodeCounter uint64
	shardingId  uint64
}

func NewUuid(redisClient *redis.RedisClient, config *config.Config) (*Uuid, error) {
	redisCounter, err := GetCounter(redisClient)
	if err != nil {
		return nil, err
	}

	baseCounter := redisCounter % config.Uuid.TotalReplicateNodes
	maxCountInNode := 2048 / config.Uuid.TotalReplicateNodes

	min := baseCounter * maxCountInNode
	max := min + maxCountInNode - 1

	return &Uuid{
		min,
		max,
		min,
		config.Uuid.ShardingId,
	}, nil
}

func GetCounter(redisClient *redis.RedisClient) (uint64, error) {
	counter := uint64(0)

	strCounter, err := redisClient.GetValue(context.Background(), COUNTER_KEY)
	if err != nil {
		err := redisClient.SetValue(context.Background(), COUNTER_KEY, "0")
		if err != nil {
			return 0, err
		}
	}

	counter, err = strconv.ParseUint(strCounter, 10, 32)
	if err != nil {
		counter = 0
		err := redisClient.SetValue(context.Background(), COUNTER_KEY, "0")
		if err != nil {
			return 0, err
		}
	}

	if counter > COUNTER_MAX {
		err := redisClient.SetValue(context.Background(), COUNTER_KEY, "0")
		if err != nil {
			return 0, err
		}
	}

	err = redisClient.IncValue(context.Background(), COUNTER_KEY)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

func (u *Uuid) incNodeCounter() uint64 {
	if u.nodeCounter >= u.max {
		u.nodeCounter = u.min
	}

	u.nodeCounter = u.nodeCounter + 1
	return u.nodeCounter
}

func (u *Uuid) GenerateUuid(ctx context.Context, uuidType constant.UuidType) uint64 {
	counterInt := u.incNodeCounter()
	now := uint64(time.Now().UnixMilli())

	uuid := now << 23
	uuid |= u.shardingId << 17
	uuid |= uint64(uuidType) << 11
	uuid |= counterInt

	logger.Info(ctx).Msgf("New uuid generated: %d. Time: %d, Sharding: %d, Type: %d, Counter: %d", uuid, now, u.shardingId, uuidType, counterInt)
	return uuid
}
