package response

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

func SendSuccess(ctx *fiber.Ctx, data interface{}) error {
	if data == nil {
		data = new(interface{})
	}
	span := trace.SpanFromContext(ctx.UserContext())
	traceId := span.SpanContext().TraceID().String()

	return ctx.Status(fiber.StatusOK).JSON(
		Messsage{
			Success:    true,
			Message:    "successfully",
			StatusCode: fiber.StatusOK,
			Data:       data,
			Trace:      traceId,
		},
	)
}
