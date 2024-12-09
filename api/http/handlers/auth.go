package handlers

import (
	presenter "NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/user"
	"NexaForm/service"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RegisterUser(authService *service.AuthService, logService *service.LoggerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		

		var req presenter.UserRegisterReq

		if err := c.BodyParser(&req); err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}
		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		u := presenter.UserRegisterToUserDomain(&req)

		newUser, err := authService.CreateUserAndSentOtp(c.Context(), u)
		fields := []zap.Field{
			//zap.String("request_id", randomString(10)),
			//zap.String("user_id", ran),
			zap.String("action", "registration"),
			//zap.String("status", randomStatus()),
			zap.String("endpoint", "api/v1/register"),
		}
		if err != nil {
			logService.LogError(c.Context(), "userService", "Generated error log", fields...)
			if errors.Is(err, user.ErrInvalidEmail) || errors.Is(err, user.ErrInvalidPassword) {
				return presenter.BadRequest(c, err)
			}
			if errors.Is(err, user.ErrEmailAlreadyExists) {
				return presenter.Conflict(c, err)
			}

			return presenter.InternalServerError(c, err)
		}

		return presenter.Created(c, "otp code successfully generated", fiber.Map{
			"user_id": newUser.ID,
		})
	}
}

// LoginUser logs in an existing user.
// @Summary Login an existing user
// @Description Authenticate a user with email and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body presenter.UserLoginReq true "User Login details"
// @Success 200 {object} map[string]interface{} "auth_token: the authentication token for the user"
// @Failure 400 {object} map[string]interface{} "error: bad request, invalid email or password"
// @Router /login [post]
func LoginUser(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.UserLoginReq

		if err := c.BodyParser(&req); err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}

		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		authToken, err := authService.Login(c.Context(), req.Email, req.Password)
		if err != nil {

			return presenter.BadRequest(c, err)
		}
		return SendUserToken(c, authToken)
	}
}
func VerifyEmail(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.EmailVerifyReq
		if err := c.BodyParser(&req); err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}
		user, err := authService.VerifyEmail(c.Context(), req.Email, req.OTP)
		if err != nil {
			return SendError(c, err, fiber.StatusBadRequest)
		}
		return presenter.OK(c, "email verified successfuly", fiber.Map{"user_id": user.ID})
	}
}
func RefreshToken(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refToken := c.GetReqHeaders()["Authorization"][0]
		if len(refToken) == 0 {
			return SendError(c, errors.New("token should be provided"), fiber.StatusBadRequest)
		}
		pureToken := strings.Split(refToken, " ")[1]
		authToken, err := authService.RefreshAuth(c.UserContext(), pureToken)
		if err != nil {

			return presenter.Unauthorized(c, err)
		}

		return SendUserToken(c, authToken)
	}
}
