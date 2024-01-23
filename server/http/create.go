package http

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/pkg/response"
)

func Create() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: response.ErrorHandler,
		})

	return app
}
