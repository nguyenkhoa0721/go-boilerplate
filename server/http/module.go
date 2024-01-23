package http

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), fx.Provide(CreateGroup))
}
