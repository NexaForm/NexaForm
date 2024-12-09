package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type Option struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	QuestionID uuid.UUID `gorm:"not null"`
	Question   Question  `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string    `gorm:"not null"`
	IsCorrect  *bool     `gorm:"null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
