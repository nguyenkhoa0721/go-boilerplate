package infra

import (
	"go-boilerplate/pkg/database/postgres"
	"go-boilerplate/pkg/database/redis"
	"go-boilerplate/pkg/otel"
	"go-boilerplate/pkg/uuid"
)

type Infra struct {
	Sql    *postgres.Sql
	Redis  *redis.RedisClient
	Uuid   *uuid.Uuid
	Tracer *otel.Tracer
}

func NewInfra(sql *postgres.Sql, uuid *uuid.Uuid, redis *redis.RedisClient, tracer *otel.Tracer) *Infra {
	return &Infra{
		Sql:    sql,
		Redis:  redis,
		Uuid:   uuid,
		Tracer: tracer,
	}
}
