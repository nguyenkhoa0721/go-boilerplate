package bootstrap

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/rs/zerolog/log"
	"go-boilerplate/config"
	"go.uber.org/fx"
	"net"
)

func Bootstrap(lc fx.Lifecycle, app *fiber.App, config *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.App.Address)
			if err != nil {
				return err
			}

			go func() {
				app.Use(
					helmet.New(),
					cors.New(
						cors.Config{
							AllowCredentials: true,
							AllowOrigins:     config.App.Cors,
						},
					),
				)

				if err := app.Listener(ln); err != nil {
					log.Error().Err(err).Msg("Server terminated unexpectedly")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Gracefully shutting down server")
			if err := app.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error occurred on server shutdown")
				return err
			}
			log.Info().Msg("Server gracefully stopped")
			return nil
		},
	})
}
