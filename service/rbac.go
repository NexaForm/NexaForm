package service

import (
	"NexaForm/internal/rbac"
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (rs *RBACService) CreateSurveyParticipants(ctx context.Context, surveyParticipants []rbac.SurveyParticipant) ([]rbac.SurveyParticipant, error) {
	return rs.surveyOps.CreateSurveyParticipants(ctx, surveyParticipants)
}
func (rs *RBACService) GetSurveyRole(ctx context.Context, id uuid.UUID) (*rbac.SurveyRole, error) {
	return rs.surveyOps.GetSurveyRole(ctx, id)
}

func (rs *RBACService) GetSurveyParticipant(ctx context.Context, id uuid.UUID) (*rbac.SurveyParticipant, error) {
	return rs.surveyOps.GetSurveyParticipant(ctx, id)
}

func (rs *RBACService) GetSurveyRolesBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyRole, error) {
	return rs.surveyOps.GetSurveyRolesBySurveyID(ctx, surveyID)
}

func (rs *RBACService) GetSurveyParticipantsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyParticipant, error) {
	logger, _ := ctx.Value("logger").(*zap.Logger)
	defer logger.Sync()
	participants, err := rs.surveyOps.GetSurveyParticipantsBySurveyID(ctx, surveyID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return participants, nil
}
