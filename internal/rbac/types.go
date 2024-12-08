package rbac

import (
	"NexaForm/internal/survey"
	"context"
	"database/sql"
	"os/user"
	"time"

	"github.com/google/uuid"
)

type Repo interface {
	// Create methods
	CreateSurveyRoles(ctx context.Context, surveyRoles []SurveyRole) ([]SurveyRole, error)
	CreateSurveyPermissions(ctx context.Context, surveyPermissions []SurveyPermission) ([]SurveyPermission, error)
	CreateSurveyParticipants(ctx context.Context, surveyParticipants []SurveyParticipant) ([]SurveyParticipant, error)

	// Get methods
	GetSurveyRole(ctx context.Context, id uuid.UUID) (*SurveyRole, error)
	GetSurveyPermission(ctx context.Context, id uuid.UUID) (*SurveyPermission, error)
	GetSurveyParticipant(ctx context.Context, id uuid.UUID) (*SurveyParticipant, error)

	// GetBySurveyID methods
	GetSurveyRolesBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyRole, error)
	GetSurveyPermissionsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyPermission, error)
	GetSurveyParticipantsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyParticipant, error)
}
type SurveyRole struct {
	ID           uuid.UUID
	SurveyID     uuid.UUID
	Survey       survey.Survey
	Name         string
	Permissions  []SurveyPermission
	Participants []SurveyParticipant
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type SurveyPermission struct {
	ID                   uuid.UUID
	SurveyID             uuid.UUID
	Name                 string
	Description          sql.NullString
	CanWatchSurvey       bool
	CanWatchExposedVotes bool
	CanVote              bool
	CanEditSurvey        bool
	CanAssignRole        bool
	CanAccessReports     bool
	SurveyRoles          []SurveyRole
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
type SurveyParticipant struct {
	ID           uuid.UUID
	SurveyID     uuid.UUID
	Survey       survey.Survey
	UserID       uuid.UUID
	User         user.User
	SurveyRoleID uuid.UUID
	SurveyRole   SurveyRole
	IsExposed    bool
	RoleExpire   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
