package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go-boilerplate/internal/user/domain"
	"go-boilerplate/pkg/response"
	"go.uber.org/fx"
)

type UserAdminHandler struct {
	fx.In

	Route       fiber.Router `name:"api-admin"`
	UserService *domain.UserService
}

func RegisterHandler(h UserAdminHandler) {
	route := h.Route.Group("/admin")
	route.Get("/check-health", h.CheckHealth)
}

func (h *UserAdminHandler) CheckHealth(ctx *fiber.Ctx) error {
	log.Info().Msg("CheckHealth From Admin Handler")
	resp := h.UserService.CheckHealth(ctx.Context())
	return response.SendSuccess(ctx, resp)
}
