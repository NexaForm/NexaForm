package entities

import (
	"time"

	"github.com/google/uuid"
)

type WalletTransaction struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SenderWalletID   uuid.UUID
	SenderWallet     Wallet `gorm:"foreignKey:SenderWalletID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReceiverWalletID uuid.UUID
	ReceiverWallet   Wallet    `gorm:"foreignKey:ReceiverWalletID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Amount           float64   `gorm:"not null"`
	TransactionTime  time.Time `gorm:"not null"`
	CreatedAt        time.Time
}
