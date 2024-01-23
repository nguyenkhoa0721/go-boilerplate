package public

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/user/domain"
	"go-boilerplate/pkg/response"
	"go.uber.org/fx"
)

type UserPublicHandler struct {
	fx.In

	Route       fiber.Router `name:"api-v1"`
	UserService *domain.UserService
}

func RegisterHandler(h UserPublicHandler) {
	route := h.Route.Group("/user")
	route.Get("/check-health", h.CheckHealth)
}

func (h *UserPublicHandler) CheckHealth(ctx *fiber.Ctx) error {
	resp := h.UserService.CheckHealth(ctx.Context())
	return response.SendSuccess(ctx, resp)
}
