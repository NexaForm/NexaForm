package survey

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

func (o *Ops) Create(ctx context.Context, user *Survey) (*Survey, error) {

	return o.repo.CreateSurvey(ctx, user)
}
func (o *Ops) GetByID(ctx context.Context, id uuid.UUID) (*Survey, error) {

	return o.repo.GetSurveyByID(ctx, id)
}
