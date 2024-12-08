package presenter

import (
	"NexaForm/internal/rbac"
	"time"

	"github.com/google/uuid"
)

type CreateRoleReq struct {
	Role []createSurveyRolesRequest `json:"role" validate:"required"`
}
type createSurveyRolesRequest struct {
	SurveyID uuid.UUID `json:"survey_id" validate:"required,uuid"`
	Name     string    `json:"name" validate:"required"`
}

type CreateSurveyPermissionsRequest []struct {
	SurveyID             uuid.UUID `json:"survey_id" validate:"required,uuid"`
	Name                 string    `json:"name" validate:"required"`
	Description          string    `json:"description,omitempty"`
	CanWatchSurvey       bool      `json:"can_watch_survey"`
	CanWatchExposedVotes bool      `json:"can_watch_exposed_votes"`
	CanVote              bool      `json:"can_vote"`
	CanEditSurvey        bool      `json:"can_edit_survey"`
	CanAssignRole        bool      `json:"can_assign_role"`
	CanAccessReports     bool      `json:"can_access_reports"`
}

type CreateSurveyParticipantsRequest []struct {
	SurveyID     uuid.UUID `json:"survey_id" validate:"required,uuid"`
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	SurveyRoleID uuid.UUID `json:"survey_role_id" validate:"required,uuid"`
	IsExposed    bool      `json:"is_exposed"`
	RoleExpire   time.Time `json:"role_expire"`
}

type SurveyRoleResponse struct {
	ID           uuid.UUID                   `json:"id"`
	SurveyID     uuid.UUID                   `json:"survey_id"`
	Name         string                      `json:"name"`
	Permissions  []SurveyPermissionResponse  `json:"permissions"`
	Participants []SurveyParticipantResponse `json:"participants"`
	CreatedAt    time.Time                   `json:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at"`
}

type SurveyPermissionResponse struct {
	ID          uuid.UUID            `json:"id"`
	SurveyID    uuid.UUID            `json:"survey_id"`
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	SurveyRoles []SurveyRoleResponse `json:"survey_roles"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type SurveyParticipantResponse struct {
	ID         uuid.UUID          `json:"id"`
	SurveyID   uuid.UUID          `json:"survey_id"`
	UserID     uuid.UUID          `json:"user_id"`
	SurveyRole SurveyRoleResponse `json:"survey_role"`
	IsExposed  bool               `json:"is_exposed"`
	RoleExpire time.Time          `json:"role_expire"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

// Conversion helpers for SurveyRole
func SurveyRoleToResponse(role *rbac.SurveyRole) *SurveyRoleResponse {
	return &SurveyRoleResponse{
		ID:           role.ID,
		SurveyID:     role.SurveyID,
		Name:         role.Name,
		Permissions:  BatchSurveyPermissionToResponse(role.Permissions),
		Participants: BatchSurveyParticipantToResponse(role.Participants),
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
	}
}

func BatchSurveyRoleToResponse(roles []rbac.SurveyRole) []SurveyRoleResponse {
	var responses []SurveyRoleResponse
	for _, role := range roles {
		responses = append(responses, *SurveyRoleToResponse(&role))
	}
	return responses
}

// Conversion helpers for SurveyPermission
func SurveyPermissionToResponse(permission *rbac.SurveyPermission) *SurveyPermissionResponse {
	return &SurveyPermissionResponse{
		ID:          permission.ID,
		SurveyID:    permission.SurveyID,
		Name:        permission.Name,
		Description: permission.Description.String,
		SurveyRoles: BatchSurveyRoleToResponse(permission.SurveyRoles),
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}

func BatchSurveyPermissionToResponse(permissions []rbac.SurveyPermission) []SurveyPermissionResponse {
	var responses []SurveyPermissionResponse
	for _, permission := range permissions {
		responses = append(responses, *SurveyPermissionToResponse(&permission))
	}
	return responses
}

// Conversion helpers for SurveyParticipant
func SurveyParticipantToResponse(participant *rbac.SurveyParticipant) *SurveyParticipantResponse {
	return &SurveyParticipantResponse{
		ID:         participant.ID,
		SurveyID:   participant.SurveyID,
		UserID:     participant.UserID,
		SurveyRole: *SurveyRoleToResponse(&participant.SurveyRole),
		IsExposed:  participant.IsExposed,
		RoleExpire: participant.RoleExpire,
		CreatedAt:  participant.CreatedAt,
		UpdatedAt:  participant.UpdatedAt,
	}
}

func BatchSurveyParticipantToResponse(participants []rbac.SurveyParticipant) []SurveyParticipantResponse {
	var responses []SurveyParticipantResponse
	for _, participant := range participants {
		responses = append(responses, *SurveyParticipantToResponse(&participant))
	}
	return responses
}
