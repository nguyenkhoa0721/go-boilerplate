package uuid

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("uuid", fx.Provide(NewUuid))
}
