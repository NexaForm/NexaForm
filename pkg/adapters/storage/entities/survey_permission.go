package entities

// import (
// 	"database/sql"
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // TODO - don't forget to change this entity for setting up your related service

// type SurveyPermission struct {
// 	ID                   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	SurveyID             uuid.UUID `gorm:"not null"`
// 	Survey               Survey    `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	Name                 string    `gorm:"not null"`
// 	Description          sql.NullString
// 	CanWatchSurvey       bool         `gorm:"not null;default:false"`
// 	CanWatchExposedVotes bool         `gorm:"not null;default:false"`
// 	CanVote              bool         `gorm:"not null;default:false"`
// 	CanEditSurvey        bool         `gorm:"not null;default:false"`
// 	CanAssignRole        bool         `gorm:"not null;default:false"`
// 	CanAccessReports     bool         `gorm:"not null;default:false"`
// 	SurveyRoles          []SurveyRole `gorm:"many2many:survey_role_permissions"`
// 	CreatedAt            time.Time
// 	UpdatedAt            time.Time
// 	DeletedAt            gorm.DeletedAt `gorm:"index"` // Soft delete support
// }
