package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO - don't forget to change this entity for setting up your related service

type SurveyPermission struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description sql.NullString
	SurveyRoles []SurveyRole `gorm:"many2many:survey_role_permissions"`
	CreatedAt   time.Time
	gorm.Model
}
