package service

import (
	"NexaForm/internal/rbac"
	"context"

	"github.com/google/uuid"
)

type RBACService struct {
	surveyOps *rbac.Ops
}

// NewSurveyService initializes SurveyService and sets up Redis client
func NewRBACService(rbacOps *rbac.Ops) *RBACService {
	return &RBACService{
		surveyOps: rbacOps,
	}
}

func (rs *RBACService) CreateSurveyRoles(ctx context.Context, surveyRoles []rbac.SurveyRole) ([]rbac.SurveyRole, error) {
	return rs.surveyOps.CreateSurveyRoles(ctx, surveyRoles)
}

func (rs *RBACService) CreateSurveyPermissions(ctx context.Context, surveyPermissions []rbac.SurveyPermission) ([]rbac.SurveyPermission, error) {
	return rs.surveyOps.CreateSurveyPermissions(ctx, surveyPermissions)
}

func (rs *RBACService) CreateSurveyParticipants(ctx context.Context, surveyParticipants []rbac.SurveyParticipant) ([]rbac.SurveyParticipant, error) {
	return rs.surveyOps.CreateSurveyParticipants(ctx, surveyParticipants)
}
func (rs *RBACService) GetSurveyRole(ctx context.Context, id uuid.UUID) (*rbac.SurveyRole, error) {
	return rs.surveyOps.GetSurveyRole(ctx, id)
}
func (rs *RBACService) GetSurveyPermission(ctx context.Context, id uuid.UUID) (*rbac.SurveyPermission, error) {
	return rs.surveyOps.GetSurveyPermission(ctx, id)
}
func (rs *RBACService) GetSurveyParticipant(ctx context.Context, id uuid.UUID) (*rbac.SurveyParticipant, error) {
	return rs.surveyOps.GetSurveyParticipant(ctx, id)
}

func (rs *RBACService) GetSurveyRolesBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyRole, error) {
	return rs.surveyOps.GetSurveyRolesBySurveyID(ctx, surveyID)
}
func (rs *RBACService) GetSurveyPermissionsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyPermission, error) {
	return rs.surveyOps.GetSurveyPermissionsBySurveyID(ctx, surveyID)
}
func (rs *RBACService) GetSurveyParticipantsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyParticipant, error) {
	return rs.surveyOps.GetSurveyParticipantsBySurveyID(ctx, surveyID)
}
