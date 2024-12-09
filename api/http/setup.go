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
	secret := []byte(cfg.Server.TokenSecret)
	api := fiberApp.Group("/api/v1")

	// register routes
	registerGlobalRoutes(api, app)
	registerSurveyRoutes(api, app, secret)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/register", handlers.RegisterUser(app.AuthService()))
	router.Post("/register/verify", handlers.VerifyEmail(app.AuthService()))
	router.Get("/refresh", handlers.RefreshToken(app.AuthService()))
	router.Post("/login", handlers.LoginUser(app.AuthService()))
	// router.Post("/survey", handlers.AddSurveyHandler(app.SurveyService()))
	// router.Post("/survey/presigned-urls", handlers.GetPresignedURLsHandler(app.SurveyService(), app.FileService()))
	// router.Get("/survey/download-urls", handlers.GetPresignedDownloadURLsHandler(app.FileService()))
	// router.Post("/survey/answer", handlers.CreateAnswerHandler(app.SurveyService()))

}
func registerSurveyRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	router.Post("/survey",
		middlewares.Auth(secret),
		handlers.AddSurveyHandler(app.SurveyService()))
	router.Post("/survey/presigned-urls",
		middlewares.Auth(secret),
		handlers.GetPresignedURLsHandler(app.SurveyService(), app.FileService()))
	router.Get("/survey/download-urls",
		middlewares.Auth(secret),
		handlers.GetPresignedDownloadURLsHandler(app.FileService()))
	router.Post("/survey/answer",
		middlewares.Auth(secret),
		handlers.CreateAnswerHandler(app.SurveyService()))
}
