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
	registerRBACRoutes(api, app, secret)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}

func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	global := router.Group("/auth", middlewares.LoggerMiddleware(app, service.ServiceAuth))

	global.Post("/register", handlers.RegisterUser(app.AuthService()))
	global.Post("/register/verify", handlers.VerifyEmail(app.AuthService()))
	global.Get("/refresh", handlers.RefreshToken(app.AuthService()))
	global.Post("/login", handlers.LoginUser(app.AuthService()))
}
func registerSurveyRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	survey := router.Group("/survey", middlewares.LoggerMiddleware(app, service.ServiceSurvey))

	survey.Post("",
		middlewares.Auth(secret),
		handlers.AddSurveyHandler(app.SurveyService()))
	survey.Post("/presigned-urls",
		middlewares.Auth(secret),
		handlers.GetPresignedURLsHandler(app.SurveyService(), app.FileService()))
	survey.Get("/download-urls",
		middlewares.Auth(secret),
		handlers.GetPresignedDownloadURLsHandler(app.FileService()))
	survey.Post("/answer",
		middlewares.Auth(secret),
		handlers.CreateAnswerHandler(app.SurveyService()))
}
func registerRBACRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	rbac := router.Group("/rbac", middlewares.LoggerMiddleware(app, service.ServiceRBAC))

	// Endpoint to create survey roles
	rbac.Post("/roles",
		middlewares.Auth(secret),
		handlers.CreateSurveyRoleHandler(app.RBACService()))

	// Endpoint to get survey roles by ID
	rbac.Get("/roles/:id",
		middlewares.Auth(secret),
		handlers.GetSurveyRoleHandler(app.RBACService()))

	// Endpoint to get survey roles by survey ID
	rbac.Get("/roles/by-survey/:survey_id",
		middlewares.Auth(secret),
		handlers.GetSurveyRolesBySurveyIDHandler(app.RBACService()))

	// Endpoint to create survey participants
	rbac.Post("/participants",
		middlewares.Auth(secret),
		handlers.CreateSurveyParticipantHandler(app.RBACService()))

	// Endpoint to get survey participants by survey ID
	rbac.Get("/participants/by-survey/:survey_id",
		middlewares.Auth(secret),
		handlers.GetSurveyParticipantsBySurveyIDHandler(app.RBACService()))
}
