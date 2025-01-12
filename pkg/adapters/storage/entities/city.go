package entities

import (
	"time"

	"github.com/google/uuid"
)

// TODO - don't forget to change this entity for setting up your related service

type City struct {
	ID               uuid.UUID     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name             string        `gorm:"unique;not null"`
	Users            []User        `gorm:"foreignKey:CityID"`
	AllowedInSurveys []AllowedCity `gorm:"foreignKey:CityID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
