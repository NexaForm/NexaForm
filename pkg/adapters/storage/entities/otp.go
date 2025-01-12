package entities

import (
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`                 // Foreign key to User
	OTPCode   string    `gorm:"type:varchar(6);not null"` // 6-digit OTP code
	OTPExpiry time.Time `gorm:"not null"`                 // Expiration timestamp
	CreatedAt time.Time `gorm:"autoCreateTime"`           // Automatically set creation time
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
