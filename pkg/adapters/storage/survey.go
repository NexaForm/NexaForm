package storage

import (
	"NexaForm/internal/survey"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type surveyRepo struct {
	db *gorm.DB
}

// NewSurveyRepo creates a new instance of the survey repository
func NewSurveyRepo(db *gorm.DB) survey.Repo {
	return &surveyRepo{
		db: db,
	}
}

// CreateSurvey inserts a new survey into the database
func (r *surveyRepo) CreateSurvey(ctx context.Context, survey *survey.Survey) (*survey.Survey, error) {
	entity := mappers.SurveyDomainToEntity(survey)
	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return nil, err
	}
	createdSurvey := mappers.SurveyEntityToDomain(entity)
	return createdSurvey, nil
}

// GetSurveyByID retrieves a survey by its ID from the database
func (r *surveyRepo) GetSurveyByID(ctx context.Context, id uuid.UUID) (*survey.Survey, error) {
	var entity entities.Survey
	err := r.db.WithContext(ctx).Preload("Questions.Options").Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mappers.SurveyEntityToDomain(&entity), nil
}
