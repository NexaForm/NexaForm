package storage

import (
	"NexaForm/internal/wallet"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"NexaForm/pkg/utils"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{
		db: db,
	}
}

func (r *walletRepo) Create(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	newWallet := mappers.WalletDomainToEntity(wallet)
	err := r.db.WithContext(ctx).Create(&newWallet).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, utils.DbErrDuplicateKey
		}
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(newWallet)
	return createdWallet, nil
}
func (r *walletRepo) GetByID(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	var w entities.Wallet
	err := r.db.WithContext(ctx).Model(&wallet.Wallet{}).Where("id=?", id).First(&w).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.WalletEntityToDomain(&w), nil
}
func (r *walletRepo) GetBalance(ctx context.Context, id uuid.UUID) (float64, error) {
	var w entities.Wallet
	if err := r.db.WithContext(ctx).Model(&wallet.Wallet{}).Where("id=?", id).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return w.Balance, nil
}
func (r *walletRepo) CreditWallet(ctx context.Context, id uuid.UUID, amount float64) error {
	var w entities.Wallet
	if err := r.db.WithContext(ctx).Model(&wallet.Wallet{}).Where("id=?", id).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	w.Balance += amount
	return r.db.WithContext(ctx).Save(&w).Error
}
func (r *walletRepo) TransferFunds(ctx context.Context, amount float64, senderID uuid.UUID, receiverID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var sender, receiver entities.Wallet
	err := tx.Model(&wallet.Wallet{}).Where("id=?", senderID).First(&sender).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil
		}
		tx.Rollback()
		return err
	}
	err = tx.Model(&wallet.Wallet{}).Where("id=?", receiverID).First(&receiver).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil
		}
		tx.Rollback()
		return err
	}
	sender.Balance -= amount
	receiver.Balance += amount
	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&receiver).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *walletRepo) GetTransactionHistory(ctx context.Context, id uuid.UUID, pageIndex int, pageSize int) ([]entities.WalletTransaction, error) {
	var w entities.Wallet
	if err := r.db.WithContext(ctx).Model(&wallet.Wallet{}).Where("id=?", id).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var transactions []entities.WalletTransaction
	err := r.db.WithContext(ctx).Model(&entities.WalletTransaction{}).Where("wallet_id=?", id).Order("created_at desc").Limit(pageSize).Offset(pageSize * pageIndex).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
