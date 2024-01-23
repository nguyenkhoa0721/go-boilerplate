package auth

import (
	"go-boilerplate/internal/auth/domain"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("auth", domain.Module())
}
