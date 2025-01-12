package entities

import (
	"NexaForm/internal/survey"
	"NexaForm/internal/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Survey struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	OwnerID     uuid.UUID `gorm:"not null"`
	Owner       User      `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string    `gorm:"not null"`
	Description *string   `gorm:"null"`
	// customization options
	Visibility    survey.VisibilityType `gorm:"type:visibility_enum;not null"`
	AllowedMinAge *int                  `gorm:"null"`
	AllowedMaxAge *int                  `gorm:"null"`
	AllowedCities []AllowedCity         `gorm:"foreignKey:SurveyID"`
	AllowedGender *user.GenderType      `gorm:"type:gender_enum;null"` // Reference PostgreSQL ENUM type
	StartTime     time.Time             `gorm:"not null"`
	EndTime       time.Time             `gorm:"not null"`
	MaxEditTime   time.Time             `gorm:"not null"`
	// end customization options
	IsOrdered          bool                `gorm:"not null"`
	IsReversable       bool                `gorm:"not null"`
	ParticipationCount int                 `gorm:"not null"`
	MaxTries           int                 `gorm:"not null"`
	Questions          []Question          `gorm:"foreignKey:SurveyID;saveAssociation:true"`
	Participants       []SurveyParticipant `gorm:"foreignKey:SurveyID"`
	SurveyRoles        []SurveyRole        `gorm:"foreignKey:SurveyID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"` // Soft delete support

}
