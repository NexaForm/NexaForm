package entities

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description string
	Roles       []Role `gorm:"many2many:role_permissions"`
	CreatedAt   time.Time
}
