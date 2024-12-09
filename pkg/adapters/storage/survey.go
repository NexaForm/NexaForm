package storage

import (
	"NexaForm/internal/survey"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"context"
	"fmt"

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

// GetQuestionsBySurveyID retrieves questions for a specific survey from the database
func (r *surveyRepo) GetQuestionsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]survey.Question, error) {
	var questionEntities []entities.Question

	err := r.db.WithContext(ctx).
		Where("survey_id = ?", surveyID).
		Preload("Attachments"). // Preload the attachments for each question
		Find(&questionEntities).Error
	if err != nil {
		return nil, err
	}

	return mappers.BatchQuestionEntityToDomain(questionEntities), nil
}

// CreateAttachments creates attachment records for specified question IDs
func (r *surveyRepo) CreateAttachments(ctx context.Context, attachments ...survey.Attachment) error {
	if len(attachments) == 0 {
		return nil // No attachments to process
	}

	// Use the BatchAttachmentDomainToEntity mapper to convert survey domain attachments to entities
	attachmentEntities := mappers.BatchAttachmentDomainToEntity(attachments)

	// Perform a bulk insert or individual inserts using the database connection
	if err := r.db.Create(&attachmentEntities).Error; err != nil {
		return err
	}

	return nil
}

// UpdateAttachments updates attachment records based on their file paths and question IDs
func (r *surveyRepo) UpdateAttachments(ctx context.Context, attachments ...survey.Attachment) error {
	if len(attachments) == 0 {
		return nil // No attachments to update
	}

	for _, attachment := range attachments {
		// Find the existing attachment by its file path and question ID
		var existingAttachment entities.Attachment
		err := r.db.WithContext(ctx).Where("file_path = ? AND question_id = ?", attachment.FilePath, attachment.QuestionID).First(&existingAttachment).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("attachment with file path %s and question ID %s not found", attachment.FilePath, attachment.QuestionID)
			}
			return err
		}

		// Update fields in the existing attachment
		existingAttachment.IsPersisted = attachment.IsPersisted

		// Save the changes
		if err := r.db.WithContext(ctx).Save(&existingAttachment).Error; err != nil {
			return err
		}
	}

	return nil
}
func (r *surveyRepo) CreateAnswer(ctx context.Context, answer survey.Answer) (*survey.Answer, error) {
	// Convert the domain Answer to an entity using the mapper
	answerEntity := mappers.AnswerDomainToEntity(&answer)

	// Insert the answer entity into the database
	err := r.db.WithContext(ctx).Create(&answerEntity).Error
	if err != nil {
		return nil, err
	}

	// No need to re-retrieve the entity; GORM populates the struct with auto-generated fields
	createdAnswer := mappers.AnswerEntityToDomain(answerEntity)

	return createdAnswer, nil
}
func (r *surveyRepo) CheckAnswerExists(ctx context.Context, questionID, userID uuid.UUID) (*survey.Answer, error) {
	var existingAnswer entities.Answer

	// Query the database to check for existing answers
	err := r.db.WithContext(ctx).
		Where("question_id = ? AND user_id = ?", questionID, userID).
		First(&existingAnswer).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No answer exists
		}
		return nil, err
	}

	// Map the entity to the domain model
	return mappers.AnswerEntityToDomain(&existingAnswer), nil
}
func (r *surveyRepo) GetSurveyByQuestionID(ctx context.Context, questionID uuid.UUID) (*survey.Survey, error) {
	var surveyEntity entities.Survey

	// Fetch the survey by joining with the questions table
	err := r.db.WithContext(ctx).
		Table("surveys").
		Joins("JOIN questions ON questions.survey_id = surveys.id").
		Where("questions.id = ?", questionID).
		First(&surveyEntity).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("survey not found for question %s", questionID.String())
		}
		return nil, err
	}

	// Explicitly preload the Questions relationship
	err = r.db.WithContext(ctx).
		Preload("Questions.Options").
		Preload("Questions.Attachments").
		Where("id = ?", surveyEntity.ID).
		First(&surveyEntity).
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to load questions for survey %s: %w", surveyEntity.ID.String(), err)
	}

	// Map the survey entity to the domain model
	return mappers.SurveyEntityToDomain(&surveyEntity), nil
}
func (r *surveyRepo) GetAnsweredQuestionsByUser(ctx context.Context, surveyID, userID uuid.UUID) ([]survey.Question, error) {
	var answeredEntities []entities.Answer

	// Query the database to fetch answers for the given survey and user using a JOIN
	err := r.db.WithContext(ctx).
		Table("answers").
		Select("answers.*").
		Joins("JOIN questions ON answers.question_id = questions.id").
		Where("questions.survey_id = ? AND answers.user_id = ?", surveyID, userID).
		Preload("Question").
		Find(&answeredEntities).
		Error
	if err != nil {
		return nil, err
	}

	// Map the fetched answers to their associated questions
	var answeredQuestions []survey.Question
	for _, answerEntity := range answeredEntities {
		if answerEntity.Question.ID != uuid.Nil {
			answeredQuestions = append(answeredQuestions, *mappers.QuestionEntityToDomain(&answerEntity.Question))
		}
	}

	return answeredQuestions, nil
}
