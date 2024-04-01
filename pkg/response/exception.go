package response

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type Messsage struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Trace      string      `json:"trace"`
}

func SendValidationError(ctx *fiber.Ctx, errs map[string]string) error {
	logger.Error(ctx.UserContext()).Msgf("validation error: %s", errs)
	return ctx.Status(fiber.StatusBadRequest).JSON(
		Messsage{
			Success:    false,
			Message:    "Validation Error",
			StatusCode: fiber.StatusBadRequest,
			Data:       errs,
		})
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.Error(ctx.UserContext()).Err(err)

	code := fiber.StatusInternalServerError
	span := trace.SpanFromContext(ctx.UserContext())
	traceId := span.SpanContext().TraceID().String()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    "Internal Server Error",
				StatusCode: code,
				Trace:      traceId,
			})
	} else if code == fiber.StatusUnauthorized {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    "Unauthorized",
				StatusCode: code,
				Trace:      traceId,
			})
	}

	return ctx.Status(code).JSON(
		Messsage{
			Success:    false,
			Message:    err.Error(),
			StatusCode: code,
			Trace:      traceId,
		})
}
