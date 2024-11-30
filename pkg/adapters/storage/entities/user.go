package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenderType defines valid gender values
type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

type User struct {
	ID              uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Primary Key
	Email           string         `gorm:"unique;not null"`                                // Unique, Non-Nullable
	IsEmailVerified *bool          `gorm:"unique;null"`
	Password        string         `gorm:"not null"`              // Non-Nullable
	FullName        *string        `gorm:"null"`                  // Nullable
	Gender          *GenderType    `gorm:"type:gender_enum;null"` // Reference PostgreSQL ENUM type
	NationalID      string         `gorm:"unique;not null"`       // Unique, Non-Nullable
	Birthday        sql.NullTime   `gorm:"null"`                  // Nullable
	RoleID          uuid.UUID      `gorm:"not null"`              // Foreign Key to Role
	Role            Role           `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	MaxSurveyCount  sql.NullInt64  `gorm:"null"`           // Nullable
	CityID          sql.NullString `gorm:"type:uuid;null"` // Nullable UUID
	City            City           `gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Surveys         []Survey       `gorm:"foreignKey:OwnerID"`
	Notifications   []Notification
	Wallet          Wallet   `gorm:"foreignKey:UserID"`
	Answers         []Answer `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"` // Soft delete support
}
