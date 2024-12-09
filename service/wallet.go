package service

import (
	"NexaForm/internal/user"
	"NexaForm/internal/wallet"
	"NexaForm/pkg/adapters/storage/entities"
	"context"

	"github.com/google/uuid"
)

type WalletService struct {
	userOps   *user.Ops
	walletOps *wallet.Ops
}

func NewWalletService(userOps *user.Ops, walletOps *wallet.Ops) *WalletService {
	return &WalletService{
		userOps:   userOps,
		walletOps: walletOps,
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return s.walletOps.Create(ctx, wallet)
}

func (s *WalletService) GetByUserId(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	return s.walletOps.GetByUserID(ctx, id)
}

func (s *WalletService) GetBalance(ctx context.Context, userID uuid.UUID) (*wallet.Wallet, error) {
	_, err := s.userOps.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.walletOps.GetBalance(ctx, userID)
}

func (s *WalletService) Deposit(ctx context.Context, id uuid.UUID, amount float64) error {
	return s.walletOps.Deposit(ctx, id, amount)
}

func (s *WalletService) Transfer(ctx context.Context, sender, receiver uuid.UUID, amount float64) error {
	return s.walletOps.TransferFunds(ctx, sender, receiver, amount)
}

func (s *WalletService) GetTransactionHistory(ctx context.Context, id uuid.UUID, pageIndex int, pageSize int) ([]entities.WalletTransaction, error) {
	return s.walletOps.GetTransactionHistory(ctx, id, pageIndex, pageSize)
}
