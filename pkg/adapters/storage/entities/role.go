package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type Role struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description string
	Users       []User       `gorm:"foreignKey:RoleID"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
	CreatedAt   time.Time
}
