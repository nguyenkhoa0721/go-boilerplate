package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Messsage struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func SendValidationError(ctx *fiber.Ctx, errs map[string]string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(
		Messsage{
			Success:    false,
			Message:    "Validation Error",
			StatusCode: fiber.StatusBadRequest,
			Data:       errs,
		})
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	logrus.Error(err)
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    "Internal Server Error",
				StatusCode: code,
			})
	} else if code == fiber.StatusUnauthorized {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    "Unauthorized",
				StatusCode: code,
			})
	}

	return ctx.Status(code).JSON(
		Messsage{
			Success:    false,
			Message:    err.Error(),
			StatusCode: code,
		})
}
