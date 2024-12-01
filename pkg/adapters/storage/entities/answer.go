package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type Answer struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	QuestionID       uuid.UUID `gorm:"not null"`
	Question         Question  `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID           uuid.UUID `gorm:"not null"`
	User             User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnswerText       string    `gorm:"type:text"` // For open-ended responses
	SelectedOptionID *uuid.UUID
	SelectedOption   *Option `gorm:"foreignKey:SelectedOptionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Nullable for open-ended answers
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
