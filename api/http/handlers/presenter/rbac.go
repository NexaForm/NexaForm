package presenter

import (
	"NexaForm/internal/rbac"
	"time"

	"github.com/google/uuid"
)

type CreateRoleReq struct {
	Role []createSurveyRolesRequest `json:"role" validate:"required"`
}
type CreateParticipantReq struct {
	Participant []createSurveyParticipantsRequest `json:"participant" validate:"required"`
}
type createSurveyRolesRequest struct {
	SurveyID             uuid.UUID `json:"survey_id" validate:"required,uuid"`
	Name                 string    `json:"name" validate:"required"`
	CanWatchSurvey       bool      `json:"can_watch_survey"`
	CanWatchExposedVotes bool      `json:"can_watch_exposed_votes"`
	CanVote              bool      `json:"can_vote"`
	CanEditSurvey        bool      `json:"can_edit_survey"`
	CanAssignRole        bool      `json:"can_assign_role"`
	CanAccessReports     bool      `json:"can_access_reports"`
}

type createSurveyParticipantsRequest struct {
	SurveyID     uuid.UUID `json:"survey_id" validate:"required,uuid"`
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	SurveyRoleID uuid.UUID `json:"survey_role_id" validate:"required,uuid"`
	IsExposed    bool      `json:"is_exposed"`
	RoleExpire   time.Time `json:"role_expire"`
}

type SurveyRoleResponse struct {
	ID                   uuid.UUID `json:"id"`
	SurveyID             uuid.UUID `json:"survey_id"`
	Name                 string    `json:"name"`
	CanWatchSurvey       bool      `json:"can_watch_survey"`
	CanWatchExposedVotes bool      `json:"can_watch_exposed_votes"`
	CanVote              bool      `json:"can_vote"`
	CanEditSurvey        bool      `json:"can_edit_survey"`
	CanAssignRole        bool      `json:"can_assign_role"`
	CanAccessReports     bool      `json:"can_access_reports"`
}

type SurveyParticipantResponse struct {
	ID         uuid.UUID          `json:"id"`
	SurveyID   uuid.UUID          `json:"survey_id"`
	UserID     uuid.UUID          `json:"user_id"`
	SurveyRole SurveyRoleResponse `json:"survey_role"`
	IsExposed  bool               `json:"is_exposed"`
	RoleExpire time.Time          `json:"role_expire"`
}

// Conversion helpers for SurveyRole
func SurveyRoleToResponse(role *rbac.SurveyRole) *SurveyRoleResponse {
	return &SurveyRoleResponse{
		ID:                   role.ID,
		SurveyID:             role.SurveyID,
		Name:                 role.Name,
		CanWatchSurvey:       role.CanWatchSurvey,
		CanWatchExposedVotes: role.CanWatchExposedVotes,
		CanVote:              role.CanVote,
		CanEditSurvey:        role.CanEditSurvey,
		CanAssignRole:        role.CanAssignRole,
		CanAccessReports:     role.CanAccessReports,
	}
}

func BatchSurveyRoleToResponse(roles []rbac.SurveyRole) []SurveyRoleResponse {
	var responses []SurveyRoleResponse
	for _, role := range roles {
		responses = append(responses, *SurveyRoleToResponse(&role))
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
	}
}

func BatchSurveyParticipantToResponse(participants []rbac.SurveyParticipant) []SurveyParticipantResponse {
	var responses []SurveyParticipantResponse
	for _, participant := range participants {
		responses = append(responses, *SurveyParticipantToResponse(&participant))
	}
	return responses
}
