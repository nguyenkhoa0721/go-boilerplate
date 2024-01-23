package module

import (
	"go-boilerplate/config"
	"go-boilerplate/internal/auth"
	"go-boilerplate/internal/user"
	"go-boilerplate/pkg/database/postgres"
	"go-boilerplate/pkg/database/redis"
	"go-boilerplate/pkg/infra"
	"go-boilerplate/pkg/uuid"
	"go-boilerplate/server"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("fxmodule",
		config.Module(),
		server.Module(),
		postgres.Module(),
		redis.Module(),
		uuid.Module(),
		infra.Module(),
		user.Module(),
		auth.Module(),
	)
}
