package http

import (
	"NexaForm/api/http/handlers"
	"NexaForm/config"
	"NexaForm/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg config.Config, app *service.AppContainer) {
	fiberApp := fiber.New(fiber.Config{})

	api := fiberApp.Group("/api/v1")

	// register global routes
	registerGlobalRoutes(api, app)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/register", handlers.RegisterUser(app.AuthService()))
	router.Post("/register/verify", handlers.VerifyEmail(app.AuthService()))
	router.Get("/refresh", handlers.RefreshToken(app.AuthService()))
	router.Post("/login", handlers.LoginUser(app.AuthService()))
}
