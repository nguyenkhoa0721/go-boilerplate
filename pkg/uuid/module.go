package uuid

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("module", fx.Provide(NewUuid))
}
