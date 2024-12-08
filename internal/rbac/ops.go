package rbac

import (
	"context"

	"github.com/google/uuid"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}
func (o *Ops) CreateSurveyRoles(ctx context.Context, surveyRoles []SurveyRole) ([]SurveyRole, error) {
	return o.repo.CreateSurveyRoles(ctx, surveyRoles)
}

func (o *Ops) CreateSurveyPermissions(ctx context.Context, surveyPermissions []SurveyPermission) ([]SurveyPermission, error) {
	return o.repo.CreateSurveyPermissions(ctx, surveyPermissions)
}

func (o *Ops) CreateSurveyParticipants(ctx context.Context, surveyParticipants []SurveyParticipant) ([]SurveyParticipant, error) {
	return o.repo.CreateSurveyParticipants(ctx, surveyParticipants)
}

func (o *Ops) GetSurveyRole(ctx context.Context, id uuid.UUID) (*SurveyRole, error) {
	return o.repo.GetSurveyRole(ctx, id)
}
func (o *Ops) GetSurveyPermission(ctx context.Context, id uuid.UUID) (*SurveyPermission, error) {
	return o.repo.GetSurveyPermission(ctx, id)
}
func (o *Ops) GetSurveyParticipant(ctx context.Context, id uuid.UUID) (*SurveyParticipant, error) {
	return o.repo.GetSurveyParticipant(ctx, id)
}

func (o *Ops) GetSurveyRolesBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyRole, error) {
	return o.repo.GetSurveyRolesBySurveyID(ctx, surveyID)
}
func (o *Ops) GetSurveyPermissionsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyPermission, error) {
	return o.repo.GetSurveyPermissionsBySurveyID(ctx, surveyID)
}
func (o *Ops) GetSurveyParticipantsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]SurveyParticipant, error) {
	return o.repo.GetSurveyParticipantsBySurveyID(ctx, surveyID)
}
