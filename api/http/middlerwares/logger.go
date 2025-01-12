package middlewares

import (
	"NexaForm/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// LoggerMiddleware assigns a service-specific logger to the context
func LoggerMiddleware(app *service.AppContainer, serviceName string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		reqID := uuid.New().String()

		// Attach logger for the specific service to the user context
		ctx := app.LoggerService().AttachLoggerToContext(c.UserContext(), serviceName, c, reqID)

		// Update Fiber's user context
		c.SetUserContext(ctx)

		// Continue to the next middleware/handler
		return c.Next()
	}
}
