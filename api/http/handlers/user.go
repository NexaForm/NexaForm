package handlers

import (
	"NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/user"
	"NexaForm/pkg/jwt"
	"NexaForm/service"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetAllVerifiedUsers(userService *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return SendError(c, errWrongClaimType, fiber.StatusBadRequest)
		}
		//query parameter
		page, pageSize := PageAndPageSize(c)

		users, total, err := userService.GetAllVerifiedUsers(c.UserContext(), userClaims.UserID, uint(page), uint(pageSize))
		if err != nil {
			status := fiber.StatusInternalServerError
			if errors.Is(err, user.ErrUserNotFound) {
				status = fiber.StatusBadRequest
			}
			return SendError(c, err, status)
		}
		data := presenter.NewPagination(
			presenter.BatchUsersToUserGet(users),
			uint(page),
			uint(pageSize),
			total,
		)
		return presenter.OK(c, "users successfully fetched.", data)
	}
}