package entities

import (
	"time"

	"github.com/google/uuid"
)

type Survey struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	OwnerID            uuid.UUID `gorm:"not null"`
	Owner              User      `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title              string    `gorm:"not null"`
	Description        string
	Visibility         string
	AllowedGender      string
	AllowedMinAge      *int
	AllowedMaxAge      *int
	IsOrdered          bool
	IsReversable       bool
	ParticipationCount int
	MaxTries           int
	IsFinished         bool
	StartTime          time.Time
	EndTime            time.Time
	Questions          []Question          `gorm:"foreignKey:SurveyID"`
	AllowedCities      []AllowedCity       `gorm:"foreignKey:SurveyID"`
	Participants       []SurveyParticipant `gorm:"foreignKey:SurveyID"`
	SurveyRoles        []SurveyRole        `gorm:"foreignKey:SurveyID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
