package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type SurveyRole struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID             uuid.UUID `gorm:"not null"`
	Survey               Survey    `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name                 string    `gorm:"not null"`
	CanWatchSurvey       bool      `gorm:"not null;default:false"`
	CanWatchExposedVotes bool      `gorm:"not null;default:false"`
	CanVote              bool      `gorm:"not null;default:false"`
	CanEditSurvey        bool      `gorm:"not null;default:false"`
	CanAssignRole        bool      `gorm:"not null;default:false"`
	CanAccessReports     bool      `gorm:"not null;default:false"`
	// Permissions          []SurveyPermission  `gorm:"many2many:survey_role_permissions"`
	Participants []SurveyParticipant `gorm:"foreignKey:SurveyRoleID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
