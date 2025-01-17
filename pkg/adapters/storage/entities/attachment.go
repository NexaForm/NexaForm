package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO - don't forget to change this entity for setting up your related service

type Attachment struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	QuestionID  uuid.UUID `gorm:"not null"`
	Question    Question  `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilePath    string    `gorm:"not null"`
	IsPersisted bool      `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete support

}
