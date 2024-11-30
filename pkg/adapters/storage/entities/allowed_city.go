package entities

import (
	"time"

	"github.com/google/uuid"
)

type AllowedCity struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SurveyID  uuid.UUID `gorm:"not null"`
	Survey    Survey    `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CityID    uuid.UUID `gorm:"not null"`
	City      City      `gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
