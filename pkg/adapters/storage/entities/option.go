package entities

import (
	"time"

	"github.com/google/uuid"
)

type Option struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	QuestionID uuid.UUID `gorm:"not null"`
	Question   Question  `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string    `gorm:"not null"`
	IsCorrect  *bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
