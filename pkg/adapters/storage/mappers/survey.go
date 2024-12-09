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
		ID:               entity.ID,
		SurveyID:         entity.SurveyID,
		Description:      entity.Description,
		Type:             entity.Type,
		Order:            entity.Order,
		IsConditional:    entity.IsConditional,
		TargetQuestionID: entity.TargetQuestionID,
		Options:          BatchOptionEntityToDomain(entity.Options),
		Attachments:      BatchAttachmentEntityToDomain(entity.Attachments),
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}
func questionEntityToDomain(entity entities.Question) survey.Question {
	return survey.Question{
		ID:               entity.ID,
		SurveyID:         entity.SurveyID,
		Description:      entity.Description,
		Type:             entity.Type,
		Order:            entity.Order,
		IsConditional:    entity.IsConditional,
		TargetQuestionID: entity.TargetQuestionID,
		Options:          BatchOptionEntityToDomain(entity.Options),
		Attachments:      BatchAttachmentEntityToDomain(entity.Attachments),
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}
func BatchQuestionEntityToDomain(entities []entities.Question) []survey.Question {
	return fp.Map(entities, questionEntityToDomain)
}
func QuestionDomainToEntity(domain *survey.Question) *entities.Question {
	return &entities.Question{
		SurveyID:         domain.SurveyID,
		Description:      domain.Description,
		IsConditional:    domain.IsConditional,
		TargetQuestionID: domain.TargetQuestionID,
		Options:          BatchOptionDomainToEntity(domain.Options),
		Attachments:      BatchAttachmentDomainToEntity(domain.Attachments),
		Type:             domain.Type,
		Order:            domain.Order,
	}
}
func questionDomainToEntity(domain survey.Question) entities.Question {
	return entities.Question{
		SurveyID:         domain.SurveyID,
		Description:      domain.Description,
		Type:             domain.Type,
		Order:            domain.Order,
		IsConditional:    domain.IsConditional,
		TargetQuestionID: domain.TargetQuestionID,
		Options:          BatchOptionDomainToEntity(domain.Options),
		Attachments:      BatchAttachmentDomainToEntity(domain.Attachments),
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

// attachment mapper
func AttachmentEntityToDomain(entity *entities.Attachment) *survey.Attachment {
	return &survey.Attachment{
		ID:          entity.ID,
		QuestionID:  entity.QuestionID,
		FilePath:    entity.FilePath,
		IsPersisted: entity.IsPersisted,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func AttachmentDomainToEntity(domain *survey.Attachment) *entities.Attachment {
	return &entities.Attachment{
		QuestionID: domain.QuestionID,
		FilePath:   domain.FilePath,
	}
}
func attachmentEntityToDomain(entity entities.Attachment) survey.Attachment {
	return survey.Attachment{
		ID:          entity.ID,
		QuestionID:  entity.QuestionID,
		FilePath:    entity.FilePath,
		IsPersisted: entity.IsPersisted,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func attachmentDomainToEntity(domain survey.Attachment) entities.Attachment {
	return entities.Attachment{
		QuestionID:  domain.QuestionID,
		FilePath:    domain.FilePath,
		IsPersisted: domain.IsPersisted,
	}
}

func BatchAttachmentEntityToDomain(entities []entities.Attachment) []survey.Attachment {
	return fp.Map(entities, attachmentEntityToDomain)
}

func BatchAttachmentDomainToEntity(domains []survey.Attachment) []entities.Attachment {
	return fp.Map(domains, attachmentDomainToEntity)
}

// answer mapper
func AnswerEntityToDomain(entity *entities.Answer) *survey.Answer {
	return &survey.Answer{
		ID:               entity.ID,
		QuestionID:       entity.QuestionID,
		UserID:           entity.UserID,
		AnswerText:       entity.AnswerText,
		SelectedOptionID: entity.SelectedOptionID,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}

func AnswerDomainToEntity(domain *survey.Answer) *entities.Answer {
	return &entities.Answer{
		QuestionID:       domain.QuestionID,
		UserID:           domain.UserID,
		AnswerText:       domain.AnswerText,
		SelectedOptionID: domain.SelectedOptionID,
	}
}
func answerEntityToDomain(entity entities.Answer) survey.Answer {
	return survey.Answer{
		ID:               entity.ID,
		QuestionID:       entity.QuestionID,
		UserID:           entity.UserID,
		AnswerText:       entity.AnswerText,
		SelectedOptionID: entity.SelectedOptionID,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}

func answerDomainToEntity(domain survey.Answer) entities.Answer {
	return entities.Answer{
		QuestionID:       domain.ID,
		UserID:           domain.UserID,
		AnswerText:       domain.AnswerText,
		SelectedOptionID: domain.SelectedOptionID,
	}
}

func BatchAnswerEntityToDomain(entities []entities.Answer) []survey.Answer {
	return fp.Map(entities, answerEntityToDomain)
}

func BatchAnswerDomainToEntity(domains []survey.Answer) []entities.Answer {
	return fp.Map(domains, answerDomainToEntity)
}
