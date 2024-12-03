package mappers

import (
	"NexaForm/internal/role"
	"NexaForm/internal/user"
	"NexaForm/internal/wallet"
	"NexaForm/pkg/adapters/storage/entities"
)

func WalletEntityToDomain(entity *entities.Wallet) *wallet.Wallet {
	user := user.User{
		ID:           entity.User.ID,
		FullName:     entity.User.FullName,
		Email:        entity.User.Email,
		EmailIsValid: *entity.User.IsEmailVerified,
		Password:     entity.User.Password,
		NationalID:   entity.User.NationalID,
		Role:         role.Role{ID: entity.User.Role.ID, Name: entity.User.Role.Name},
	}
	return &wallet.Wallet{
		ID:        entity.ID,
		Balance:   entity.Balance,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		UserID:    entity.UserID,
		User:      user,
	}
}
func walletEntityToDomain(entity entities.Wallet) wallet.Wallet {
	user := user.User{
		ID:           entity.User.ID,
		FullName:     entity.User.FullName,
		Email:        entity.User.Email,
		EmailIsValid: *entity.User.IsEmailVerified,
		Password:     entity.User.Password,
		NationalID:   entity.User.NationalID,
		Role:         role.Role{ID: entity.User.Role.ID, Name: entity.User.Role.Name},
	}
	return wallet.Wallet{
		ID:        entity.ID,
		Balance:   entity.Balance,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		UserID:    entity.UserID,
		User:      user,
	}
}

func WalletDomainToEntity(domainWallet *wallet.Wallet) *entities.Wallet {
	user := &entities.User{
		ID:              domainWallet.User.ID,
		FullName:        domainWallet.User.FullName,
		Email:           domainWallet.User.Email,
		IsEmailVerified: &domainWallet.User.EmailIsValid,
		Password:        domainWallet.User.Password,
		NationalID:      domainWallet.User.NationalID,
		Role:            entities.Role{ID: domainWallet.User.Role.ID, Name: domainWallet.User.Role.Name},
	}
	return &entities.Wallet{
		ID:        domainWallet.ID,
		Balance:   domainWallet.Balance,
		CreatedAt: domainWallet.CreatedAt,
		UpdatedAt: domainWallet.UpdatedAt,
		UserID:    domainWallet.User.ID,
		User:      user,
	}
}

func walletDomainToEntity(domainWallet *wallet.Wallet) entities.Wallet {
	user := &entities.User{
		ID:              domainWallet.User.ID,
		FullName:        domainWallet.User.FullName,
		Email:           domainWallet.User.Email,
		IsEmailVerified: &domainWallet.User.EmailIsValid,
		Password:        domainWallet.User.Password,
		NationalID:      domainWallet.User.NationalID,
		Role:            entities.Role{ID: domainWallet.User.Role.ID, Name: domainWallet.User.Role.Name},
	}
	return entities.Wallet{
		ID:        domainWallet.ID,
		Balance:   domainWallet.Balance,
		CreatedAt: domainWallet.CreatedAt,
		UpdatedAt: domainWallet.UpdatedAt,
		UserID:    domainWallet.User.ID,
		User:      user,
	}
}
