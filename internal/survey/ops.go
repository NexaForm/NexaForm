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
func (o *Ops) GetQuestionsBySurveyID(ctx context.Context, id uuid.UUID) ([]Question, error) {
	return o.repo.GetQuestionsBySurveyID(ctx, id)
}
func (o *Ops) CreateAttachments(ctx context.Context, attachments ...Attachment) error {
	return o.repo.CreateAttachments(ctx, attachments...)
}
func (o *Ops) UpdateAttachments(ctx context.Context, attachments ...Attachment) error {
	return o.repo.CreateAttachments(ctx, attachments...)
}
func (o *Ops) CreateAnswer(ctx context.Context, answer Answer) (*Answer, error) {
	return o.repo.CreateAnswer(ctx, answer)
}
func (o *Ops) CheckAnswerExists(ctx context.Context, questionID, userID uuid.UUID) (*Answer, error) {
	return o.repo.CheckAnswerExists(ctx, questionID, userID)
}
func (o *Ops) GetSurveyByQuestionID(ctx context.Context, questionID uuid.UUID) (*Survey, error) {
	return o.repo.GetSurveyByQuestionID(ctx, questionID)
}
func (o *Ops) GetAnsweredQuestionsByUser(ctx context.Context, surveyID, userID uuid.UUID) ([]Question, error) {
	return o.repo.GetAnsweredQuestionsByUser(ctx, surveyID, userID)
}
