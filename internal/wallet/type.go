package wallet

import (
	"NexaForm/internal/user"
	"NexaForm/pkg/adapters/storage/entities"
	"context"
	"time"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, wallet *Wallet) (*Wallet, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*Wallet, error)
	GetBalance(ctx context.Context, userId uuid.UUID) (float64, error)
	Deposit(ctx context.Context, id uuid.UUID, amount float64) error
	TransferFunds(ctx context.Context, abount float64, senderID uuid.UUID, receiverID uuid.UUID) error
	GetTransactionHistory(ctx context.Context, id uuid.UUID, pageIndex int, pageSize int) ([]entities.WalletTransaction, error)
}

type Wallet struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	User      user.User
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
