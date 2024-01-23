package public

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("auth_public_presenter")
}
