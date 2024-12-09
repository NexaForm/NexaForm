package handlers

import (
	"NexaForm/api/http/handlers/presenter"
	"NexaForm/pkg/jwt"
	"NexaForm/service"

	"github.com/gofiber/fiber/v2"
)

func GetBalence(walletService *service.WalletService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		userClaim, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return SendError(c, errWrongClaimType, fiber.StatusBadRequest)
		}
		userID := userClaim.UserID
		wallet, err := walletService.GetBalance(c.UserContext(), userID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}
		data := presenter.DomainWalletToGetWalletResponse(wallet)
		return presenter.OK(c, "wallet fetch successfully", data)
	}
}
