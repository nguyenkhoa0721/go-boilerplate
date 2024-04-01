package http

import (
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"go-boilerplate/config"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/otel"
	"go-boilerplate/pkg/response"
)

func Create(config *config.Config, tracer *otel.Tracer) *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: response.ErrorHandler,
		})

	tracer.Init()

	app.Use(
		helmet.New(),
		cors.New(
			cors.Config{
				AllowCredentials: true,
				AllowOrigins:     config.App.Cors,
			},
		),
		otelfiber.Middleware(),
		logger.HttpLogger(),
	)

	return app
}
