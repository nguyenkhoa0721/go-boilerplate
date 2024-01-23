package user

import (
	"go-boilerplate/internal/user/data"
	"go-boilerplate/internal/user/domain"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user", domain.Module(), data.Module())
}
