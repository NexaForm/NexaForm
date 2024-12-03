package wallet

import (
	"NexaForm/internal/user"
	"context"
	"time"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, wallet *Wallet) (*Wallet, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Wallet, error)
	GetBalance(ctx context.Context, id uuid.UUID) (float64, error)
	CreditWallet(ctx context.Context, id uuid.UUID, amount float64) error
	TransferFunds(ctx context.Context, abount float64, senderID uuid.UUID, receiverID uuid.UUID) error
	//return slice of transaction
	GetTransactionHistory(ctx context.Context, id uuid.UUID, pageIndex int, pageSize int) ([]float64, int, float64, error)
}

type Wallet struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	User      user.User
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
