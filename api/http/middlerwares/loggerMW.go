package middlewares

import (
	"NexaForm/service"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LoggerMiddleware(loggerService *service.LoggerService, serviceName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		ctx := context.WithValue(context.Background(), "fiber_ctx", c)

		loggerService.LogInfo(ctx, serviceName, "Incoming request",
			zap.String("method", c.Method()),
			zap.String("url", c.OriginalURL()),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		)

		err := c.Next()

		latency := time.Since(startTime)
		loggerService.LogInfo(ctx, serviceName, "Outgoin response",
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", latency),
		)

		if err != nil {
			loggerService.LogError(ctx, serviceName, "error occruing during request",
				zap.Error(err))
		}
		return err

	}
}
