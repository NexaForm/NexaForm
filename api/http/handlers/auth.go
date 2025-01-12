package handlers

import (
	presenter "NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/user"
	"NexaForm/service"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("RegisterUser")
		var req presenter.UserRegisterReq

		if err := c.BodyParser(&req); err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}
		err := BodyValidator(req)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}

		u := presenter.UserRegisterToUserDomain(&req)

		newUser, err := authService.CreateUserAndSentOtp(c.Context(), u)
		if err != nil {
			if errors.Is(err, user.ErrInvalidEmail) || errors.Is(err, user.ErrInvalidPassword) {
				logger.Error(err.Error())
				return presenter.BadRequest(c, err)
			}
			if errors.Is(err, user.ErrEmailAlreadyExists) {
				logger.Error(err.Error())
				return presenter.Conflict(c, err)
			}
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}
		logger.Info("otp code successfuly generated")
		return presenter.Created(c, "otp code successfuly generated", fiber.Map{
			"user_id": newUser.ID,
		})
	}
}

func LoginUser(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("LoginUser")
		var req presenter.UserLoginReq

		if err := c.BodyParser(&req); err != nil {
			logger.Error(err.Error())
			return SendError(c, err, fiber.StatusBadRequest)
		}

		err := BodyValidator(req)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}

		authToken, err := authService.Login(c.Context(), req.Email, req.Password)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}
		logger.Info("access: " + authToken.AuthorizationToken + "\t" + "refresh: " + authToken.RefreshToken)
		return SendUserToken(c, authToken)
	}
}
func VerifyEmail(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("VerifyEmail")
		var req presenter.EmailVerifyReq
		if err := c.BodyParser(&req); err != nil {
			logger.Error(err.Error())
			return SendError(c, err, fiber.StatusBadRequest)
		}
		user, err := authService.VerifyEmail(c.Context(), req.Email, req.OTP)
		if err != nil {
			logger.Error(err.Error())
			return SendError(c, err, fiber.StatusBadRequest)
		}
		logger.Info("email verified successfuly")
		return presenter.OK(c, "email verified successfuly", fiber.Map{"user_id": user.ID})
	}
}
func RefreshToken(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("RefreshToken")
		refToken := c.GetReqHeaders()["Authorization"][0]
		if len(refToken) == 0 {
			logger.Error("token should be provided")
			return SendError(c, errors.New("token should be provided"), fiber.StatusBadRequest)
		}
		pureToken := strings.Split(refToken, " ")[1]
		authToken, err := authService.RefreshAuth(c.UserContext(), pureToken)
		if err != nil {
			logger.Error(err.Error())
			return presenter.Unauthorized(c, err)
		}
		logger.Info("access: " + authToken.AuthorizationToken + "\t" + "refresh: " + authToken.RefreshToken)
		return SendUserToken(c, authToken)
	}
}
