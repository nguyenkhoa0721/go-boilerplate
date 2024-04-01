package email

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("email", fx.Provide(NewEmailService))
}
