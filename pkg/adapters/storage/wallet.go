package storage

import (
	"NexaForm/internal/wallet"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

func (db *walletRepo) Create(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return nil, nil
}
func (db *walletRepo) GetByID(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	return nil, nil
}
func (db *walletRepo) GetBalance(ctx context.Context, id uuid.UUID) (float64, error) {
	return 0.0, nil
}
func (db *walletRepo) CreditWallet(ctx context.Context, id uuid.UUID, amount float64) error {
	return nil
}
func (db *walletRepo) TransferFunds(ctx context.Context, senderID uuid.UUID, receiverID uuid.UUID) error {
	return nil
}
func (db *walletRepo) GetTransactionHistory(ctx context.Context, id uuid.UUID) ([]float64, error) {
	return nil, nil
}
