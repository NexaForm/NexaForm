package entities

import (
	"NexaForm/internal/survey"
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID            uuid.UUID           `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID      uuid.UUID           `gorm:"not null"`
	Survey        Survey              `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description   string              `gorm:"not null"`
	Type          survey.QuestionType `gorm:"type:question_type_enum;not null"`
	Order         int
	IsConditional bool         `gorm:"default:false"`
	Options       []Option     `gorm:"foreignKey:QuestionID;saveAssociation:true"`
	Attachments   []Attachment `gorm:"foreignKey:QuestionID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
