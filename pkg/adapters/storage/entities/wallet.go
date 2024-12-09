package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type Wallet struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"not null"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Use pointer to avoid recursion
	Balance   float64   `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
