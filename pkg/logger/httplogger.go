package logger

import (
	"github.com/gofiber/fiber/v2"
)

func HttpLogger() fiber.Handler {
	//log ip, latency, status, method, url, error, time, message, trace

	return func(ctx *fiber.Ctx) error {
		Info(ctx.UserContext()).
			Str("ip", ctx.IP()).
			Str("url", ctx.OriginalURL()).
			Str("method", ctx.Method()).
			Str("body", string(ctx.Body())).
			Str("query", string(ctx.Request().URI().QueryString())).
			Str("params", string(ctx.Request().URI().QueryArgs().String())).
			Str("user_agent", ctx.Get("User-Agent")).
			Send()

		return ctx.Next()
	}
}
