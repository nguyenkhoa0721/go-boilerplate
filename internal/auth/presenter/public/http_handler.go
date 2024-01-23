package public

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/auth/domain"
	"go-boilerplate/internal/auth/domain/validation"
	"go-boilerplate/pkg/response"
	"go-boilerplate/pkg/validator"
	"go.uber.org/fx"
)

type AuthPublicHandler struct {
	fx.In

	Route       fiber.Router `name:"api-v1"`
	AuthService *domain.AuthService
}

func RegisterHandler(h AuthPublicHandler) {
	route := h.Route.Group("/auth")
	route.Post("/login", h.Login)
	route.Post("/signup", h.Signup)
}

func (h AuthPublicHandler) Login(ctx *fiber.Ctx) error {
	body := new(validation.LoginRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	validationErrs, ok := validator.Validate(body)
	if !ok {
		return response.SendValidationError(ctx, validationErrs)
	}

	resp, err := h.AuthService.Login(body)
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, resp)
}

func (h AuthPublicHandler) Signup(ctx *fiber.Ctx) error {
	body := new(validation.SignupRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	validationErrs, ok := validator.Validate(body)
	if !ok {
		return response.SendValidationError(ctx, validationErrs)
	}

	if err := h.AuthService.Signup(body); err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}
