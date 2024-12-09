package entities

import (
	"NexaForm/internal/survey"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID               uuid.UUID           `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID         uuid.UUID           `gorm:"not null"`
	Survey           Survey              `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description      string              `gorm:"not null"`
	Type             survey.QuestionType `gorm:"type:question_type_enum;not null"`
	Order            int
	IsConditional    bool         `gorm:"default:false"`
	TargetQuestionID *uuid.UUID   `gorm:"type:uuid;default:null"`
	TargetQuestion   *Question    `gorm:"foreignKey:TargetQuestionID"`
	Options          []Option     `gorm:"foreignKey:QuestionID;saveAssociation:true"`
	Attachments      []Attachment `gorm:"foreignKey:QuestionID;saveAssociation:true"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"` // Soft delete support

}
