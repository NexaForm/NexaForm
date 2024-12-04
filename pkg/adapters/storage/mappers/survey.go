package mappers

import (
	"NexaForm/internal/survey"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/fp"
)

// survey mappers
func SurveyEntityToDomain(entity *entities.Survey) *survey.Survey {
	return &survey.Survey{
		ID:                 entity.ID,
		OwnerID:            entity.OwnerID,
		Title:              entity.Title,
		Description:        *entity.Description,
		StartTime:          entity.EndTime,
		EndTime:            entity.EndTime,
		MaxEditTime:        entity.MaxEditTime,
		Visibility:         entity.Visibility,
		AllowedMinAge:      *entity.AllowedMinAge,
		AllowedMaxAge:      *entity.AllowedMaxAge,
		AllowedGender:      *entity.AllowedGender,
		IsOrdered:          entity.IsOrdered,
		IsReversable:       entity.IsReversable,
		ParticipationCount: entity.ParticipationCount,
		MaxTries:           entity.MaxTries,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
		Questions:          BatchQuestionEntityToDomain(entity.Questions),
	}
}
func SurveyDomainToEntity(domain *survey.Survey) *entities.Survey {
	return &entities.Survey{
		OwnerID:            domain.OwnerID,
		Title:              domain.Title,
		Description:        &domain.Description,
		Visibility:         domain.Visibility,
		StartTime:          domain.StartTime,
		EndTime:            domain.EndTime,
		AllowedMinAge:      &domain.AllowedMinAge,
		AllowedMaxAge:      &domain.AllowedMaxAge,
		AllowedGender:      &domain.AllowedGender,
		MaxEditTime:        domain.MaxEditTime,
		IsOrdered:          domain.IsOrdered,
		IsReversable:       domain.IsOrdered,
		ParticipationCount: domain.ParticipationCount,
		MaxTries:           domain.MaxTries,
		Questions:          BatchQuestionDomainToEntity(domain.Questions),
	}
}

// question mappers
func QuestionEntityToDomain(entity *entities.Question) *survey.Question {
	return &survey.Question{
		ID:            entity.ID,
		SurveyID:      entity.SurveyID,
		Description:   entity.Description,
		Type:          entity.Type,
		Order:         entity.Order,
		IsConditional: entity.IsConditional,
		Options:       BatchOptionEntityToDomain(entity.Options),
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}
func questionEntityToDomain(entity entities.Question) survey.Question {
	return survey.Question{
		ID:            entity.ID,
		SurveyID:      entity.ID,
		Description:   entity.Description,
		Type:          entity.Type,
		Order:         entity.Order,
		IsConditional: entity.IsConditional,
		Options:       BatchOptionEntityToDomain(entity.Options),
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}
func BatchQuestionEntityToDomain(entities []entities.Question) []survey.Question {
	return fp.Map(entities, questionEntityToDomain)
}
func QuestionDomainToEntity(domain *survey.Question) *entities.Question {
	return &entities.Question{
		SurveyID:      domain.SurveyID,
		Description:   domain.Description,
		IsConditional: domain.IsConditional,
		Options:       BatchOptionDomainToEntity(domain.Options),
		Type:          domain.Type,
		Order:         domain.Order,
	}
}
func questionDomainToEntity(domain survey.Question) entities.Question {
	return entities.Question{
		SurveyID:      domain.SurveyID,
		Description:   domain.Description,
		Type:          domain.Type,
		Order:         domain.Order,
		IsConditional: domain.IsConditional,
		Options:       BatchOptionDomainToEntity(domain.Options),
	}
}
func BatchQuestionDomainToEntity(domains []survey.Question) []entities.Question {
	return fp.Map(domains, questionDomainToEntity)
}

// option mappers
func OptionEntityToDomain(entity *entities.Option) *survey.Option {
	return &survey.Option{
		ID:         entity.ID,
		QuestionID: entity.QuestionID,
		Text:       entity.Text,
		IsCorrect:  entity.IsCorrect,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.CreatedAt,
	}
}
func optionEntityToDomain(entity entities.Option) survey.Option {
	return survey.Option{
		ID:         entity.ID,
		QuestionID: entity.QuestionID,
		Text:       entity.Text,
		IsCorrect:  entity.IsCorrect,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.CreatedAt,
	}
}
func BatchOptionEntityToDomain(options []entities.Option) []survey.Option {
	return fp.Map(options, optionEntityToDomain)
}
func OptionDomainToEntity(domain *survey.Option) *entities.Option {
	return &entities.Option{
		QuestionID: domain.QuestionID,
		Text:       domain.Text,
		IsCorrect:  domain.IsCorrect,
	}
}
func optionDomainToEntity(domain survey.Option) entities.Option {
	return entities.Option{
		QuestionID: domain.QuestionID,
		Text:       domain.Text,
		IsCorrect:  domain.IsCorrect,
	}
}
func BatchOptionDomainToEntity(options []survey.Option) []entities.Option {
	return fp.Map(options, optionDomainToEntity)
}
