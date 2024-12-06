package wallet

import (
	"NexaForm/pkg/adapters/storage/entities"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (ops *Ops) Create(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	if wallet == nil {
		return nil, fmt.Errorf("wallet cannot be nil")
	}
	if wallet.ID == uuid.Nil {
		return nil, fmt.Errorf("invalid wallet ID")
	}
	if wallet.UserID == uuid.Nil {
		return nil, fmt.Errorf("user ID cannot be empty")
	}
	if wallet.Balance < 0 {
		return nil, fmt.Errorf("balance cannot be negative")
	}
	if wallet.CreatedAt.IsZero() {
		wallet.CreatedAt = time.Now().UTC()
	}

	return ops.repo.Create(ctx, wallet)
}
func (ops *Ops) GetByUserID(ctx context.Context, id uuid.UUID) (*Wallet, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid wallet ID")
	}
	wallet, err := ops.GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
func (ops *Ops) GetBalance(ctx context.Context, id uuid.UUID) (float64, error) {
	if id == uuid.Nil {
		return 0, fmt.Errorf("invalid wallet ID")
	}
	amount, err := ops.GetBalance(ctx, id)
	if err != nil {
		return 0, err
	}
	return amount, nil
}
func (ops *Ops) Deposit(ctx context.Context, id uuid.UUID, amount float64) error {
	if amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}
	if id == uuid.Nil {
		return fmt.Errorf("invalid wallet ID")
	}
	if err := ops.Deposit(ctx, id, amount); err != nil {
		return err
	}
	return nil
}
func (ops *Ops) TransferFunds(ctx context.Context, senderID uuid.UUID, receiverID uuid.UUID, amount float64) error {
	if senderID == uuid.Nil {
		return fmt.Errorf("invalid sender ID")
	}
	if receiverID == uuid.Nil {
		return fmt.Errorf("invalid receiver ID")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if senderID == receiverID {
		return fmt.Errorf("sender and receiver IDs cannot be the same")
	}
	return nil
}

func (ops *Ops) GetTransactionHistory(ctx context.Context, id uuid.UUID, pageIndex int, pageSize int) ([]entities.WalletTransaction, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid wallet ID")
	}
	transactions, err := ops.repo.GetTransactionHistory(ctx, id, pageIndex, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction history: %w", err)
	}
	return transactions, nil
}
