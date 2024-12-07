package service

import (
	"NexaForm/internal/survey"
	"context"

	"github.com/google/uuid"
)

type SurveyService struct {
	surveyOps *survey.Ops
}

func NewSurveyService(surveyOps *survey.Ops) *SurveyService {
	return &SurveyService{
		surveyOps: surveyOps,
	}
}

func (ss *SurveyService) CreateSurvey(ctx context.Context, survey *survey.Survey) (*survey.Survey, error) {
	return ss.surveyOps.Create(ctx, survey)
}

func (ss *SurveyService) GetSurveyByID(ctx context.Context, surveyId uuid.UUID) (*survey.Survey, error) {
	return ss.surveyOps.GetByID(ctx, surveyId)
}
