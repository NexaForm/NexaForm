package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type SurveyParticipant struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID     uuid.UUID  `gorm:"not null"`
	Survey       Survey     `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Enforce foreign key
	UserID       uuid.UUID  `gorm:"not null"`
	User         User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Enforce foreign key
	SurveyRoleID uuid.UUID  `gorm:"not null"`                                                       // Ensure role linkage
	SurveyRole   SurveyRole `gorm:"foreignKey:SurveyRoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleExpire   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
