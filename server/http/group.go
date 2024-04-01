package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type RouteGroup struct {
	fx.Out

	APIV1 fiber.Router `name:"api-v1"`
}

func CreateGroup(app *fiber.App) RouteGroup {
	return RouteGroup{
		APIV1: app.Group("/api/v1"),
	}
}
