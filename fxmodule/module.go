package fxmodule

import (
	"go-boilerplate/config"
	"go-boilerplate/pkg/database/postgres"
	"go-boilerplate/pkg/database/redis"
	"go-boilerplate/pkg/email"
	"go-boilerplate/pkg/infra"
	"go-boilerplate/pkg/otel"
	"go-boilerplate/pkg/uuid"
	"go-boilerplate/pkg/worker"
	"go-boilerplate/server"
	"go.uber.org/fx"
)

func FeatureModule() fx.Option {
	return fx.Module("feature_module")
}

func PkgModule() fx.Option {
	return fx.Module("pkg_module",
		config.Module(),
		server.Module(),
		postgres.Module(),
		redis.Module(),
		uuid.Module(),
		infra.Module(),
		worker.Module(),
		email.Module(),
		otel.Module(),
	)
}
