package http

import (
	"NexaForm/api/http/handlers"
	middlewares "NexaForm/api/http/middlerwares"
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
	secret := []byte(cfg.Server.TokenSecret)
	registerUserRoutes(api, app, secret)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/register", handlers.RegisterUser(app.AuthService()))
	router.Post("/register/verify", handlers.VerifyEmail(app.AuthService()))
	router.Get("/refresh", handlers.RefreshToken(app.AuthService()))
	router.Post("/login", handlers.LoginUser(app.AuthService()))
}

func registerUserRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	router = router.Group("/users")

	router.Get("",
		middlewares.Auth(secret),
		handlers.GetAllVerifiedUsers(app.UserService()),
	)
}
