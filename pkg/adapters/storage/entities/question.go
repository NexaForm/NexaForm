package entities

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID    uuid.UUID `gorm:"not null"`
	Survey      Survey    `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string    `gorm:"not null"`
	Type        string
	Order       int
	Options     []Option     `gorm:"foreignKey:QuestionID"`
	Attachments []Attachment `gorm:"foreignKey:QuestionID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}