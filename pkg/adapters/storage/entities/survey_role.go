package entities

import (
	"time"

	"github.com/google/uuid"
)

type SurveyRole struct {
	ID           uuid.UUID           `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID     uuid.UUID           `gorm:"not null"`
	Survey       Survey              `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         string              `gorm:"not null"`
	Permissions  []SurveyPermission  `gorm:"many2many:survey_role_permissions"`
	Participants []SurveyParticipant `gorm:"foreignKey:SurveyRoleID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
