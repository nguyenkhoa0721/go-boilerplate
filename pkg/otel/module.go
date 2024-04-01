package otel

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("tracer", fx.Provide(NewTracer))
}
