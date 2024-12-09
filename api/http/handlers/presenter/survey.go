package presenter

import (
	"NexaForm/internal/survey"
	"NexaForm/internal/user"
	"time"

	"github.com/google/uuid"
)

type AddSurveyRequest struct {
	OwnerID       string                `json:"owner_id" validate:"required,uuid"`
	Title         string                `json:"title" validate:"required"`
	Description   string                `json:"description"`
	StartTime     time.Time             `json:"start_time" validate:"required"`
	EndTime       time.Time             `json:"end_time" validate:"required"`
	MaxEditTime   time.Time             `json:"max_edit_time" validate:"required"`
	Visibility    survey.VisibilityType `json:"visibility" validate:"required"`
	AllowedMinAge int                   `json:"allowed_min_age"`
	AllowedMaxAge int                   `json:"allowed_max_age"`
	AllowedGender string                `json:"allowed_gender"`
	IsOrdered     bool                  `json:"is_ordered"`
	IsReversable  bool                  `json:"is_reversable"`
	MaxTries      int                   `json:"max_tries"`
	Questions     []AddQuestionRequest  `json:"questions" validate:"required,dive"`
}

type AddQuestionRequest struct {
	Description   string             `json:"description" validate:"required"`
	Type          string             `json:"type" validate:"required"`
	Order         int                `json:"order"`
	IsConditional bool               `json:"is_conditional"`
	Options       []AddOptionRequest `json:"options" validate:"required,dive"`
}

type AddOptionRequest struct {
	Text      string `json:"text" validate:"required"`
	IsCorrect *bool  `json:"is_correct"`
}

var FileRequests []struct {
	QuestionID  uuid.UUID `json:"question_id"`
	FileName    string    `json:"file_name"`
	ContentType string    `json:"content_type"`
}

type CreateAnswerRequest struct {
	QuestionID       uuid.UUID  `json:"question_id" validate:"required,uuid"`
	AnswerText       string     `json:"answer_text"`
	SelectedOptionID *uuid.UUID `json:"selected_option_id" validate:"required"`
}

func AddSurveyRequestToSurveyDomain(req *AddSurveyRequest) *survey.Survey {
	var questions []survey.Question
	for _, q := range req.Questions {
		var options []survey.Option
		for _, o := range q.Options {
			options = append(options, survey.Option{
				Text:      o.Text,
				IsCorrect: o.IsCorrect,
			})
		}
		questions = append(questions, survey.Question{
			Description:   q.Description,
			Type:          survey.QuestionType(q.Type),
			Order:         q.Order,
			IsConditional: q.IsConditional,
			Options:       options,
		})
	}

	return &survey.Survey{
		OwnerID:       uuid.MustParse(req.OwnerID),
		Title:         req.Title,
		Description:   req.Description,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		MaxEditTime:   req.MaxEditTime,
		Visibility:    req.Visibility,
		AllowedMinAge: req.AllowedMinAge,
		AllowedMaxAge: req.AllowedMaxAge,
		AllowedGender: user.GenderType(req.AllowedGender),
		IsOrdered:     req.IsOrdered,
		IsReversable:  req.IsReversable,
		MaxTries:      req.MaxTries,
		Questions:     questions,
	}
}
