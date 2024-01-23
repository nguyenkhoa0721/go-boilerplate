package data

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("user_domain", fx.Provide(NewUserRepo))
}
