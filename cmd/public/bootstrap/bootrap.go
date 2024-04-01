package bootstrap

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/config"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/worker"
	"go.uber.org/fx"
	"net"
)

func Bootstrap(lc fx.Lifecycle, app *fiber.App, config *config.Config, worker *worker.Worker) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.App.Address)
			if err != nil {
				return err
			}

			go func() {
				if err := app.Listener(ln); err != nil {
					logger.Error(ctx).Err(err).Msg("Server terminated unexpectedly")
				}
			}()

			go func() {
				err := worker.StartWorker()
				if err != nil {
					logger.Error(ctx).Err(err).Msg("Failed to start worker")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info(ctx).Msg("Gracefully shutting down server")
			if err := app.Shutdown(); err != nil {
				logger.Error(ctx).Err(err).Msg("error occurred on server shutdown")
				return err
			}
			logger.Info(ctx).Msg("Server gracefully stopped")
			return nil
		},
	})
}
