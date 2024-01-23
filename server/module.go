package server

import (
	"go-boilerplate/server/http"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
