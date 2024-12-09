package presenter

import (
	"NexaForm/internal/wallet"

	"github.com/google/uuid"
)

type GetBalanceResponse struct {
	ID      uuid.UUID `json:"wallet_id"`
	Balance float64   `json:"balance"`
}

func DomainWalletToGetWalletResponse(wallet *wallet.Wallet) *GetBalanceResponse {
	w := &GetBalanceResponse{
		ID:      wallet.ID,
		Balance: wallet.Balance,
	}
	return w
}
