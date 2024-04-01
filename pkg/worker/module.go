package worker

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("worker", fx.Provide(NewWorker), fx.Provide(NewTask))
}
