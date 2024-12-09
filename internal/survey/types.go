package survey

import (
	"NexaForm/internal/user"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTitleEmpty = errors.New("title cannot be empty")
	ErrExceedsMinCharacter = errors.New("title exceeds maximum length of 255 characters")
)

type VisibilityType string
type QuestionType string

const (
	POLL QuestionType = "Poll"
	QUIZ QuestionType = "Quiz"
	TEXT QuestionType = "Text"
)

const (
	ALL         VisibilityType = "All"
	OWNER_ADMIN VisibilityType = "Owner_Admin"
	NO_ONE      VisibilityType = "No_One"
)

type Repo interface {
	CreateSurvey(ctx context.Context, survey *Survey) (*Survey, error)
	GetSurveyByID(ctx context.Context, id uuid.UUID) (*Survey, error)
	GetQuestionsBySurveyID(ctx context.Context, id uuid.UUID) ([]Question, error)
	CreateAttachments(ctx context.Context, attachments ...Attachment) error
	UpdateAttachments(ctx context.Context, attachments ...Attachment) error
	CreateAnswer(ctx context.Context, answer Answer) (*Answer, error)
	CheckAnswerExists(ctx context.Context, questionID, userID uuid.UUID) (*Answer, error)
	GetSurveyByQuestionID(ctx context.Context, questionID uuid.UUID) (*Survey, error)
	GetAnsweredQuestionsByUser(ctx context.Context, surveyID, userID uuid.UUID) ([]Question, error)
}
type Survey struct {
	ID                 uuid.UUID
	OwnerID            uuid.UUID
	Title              string
	Description        string
	StartTime          time.Time
	EndTime            time.Time
	Visibility         VisibilityType
	AllowedMinAge      int // todo not negetive
	AllowedMaxAge      int //
	AllowedGender      user.GenderType
	MaxEditTime        time.Time
	IsOrdered          bool
	IsReversable       bool
	ParticipationCount int
	MaxTries           int
	Questions          []Question
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Question struct {
	ID               uuid.UUID
	SurveyID         uuid.UUID
	Description      string // to do validate
	Type             QuestionType
	Order            int // to do validate: if survey.IsOrdered=true => Order should be 1, 2, 3, 4 ..  len survey.Questions
	IsConditional    bool
	TargetQuestionID *uuid.UUID
	Options          []Option
	Attachments      []Attachment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Option struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Text       string // validate
	IsCorrect  *bool 
	CreatedAt  time.Time 
	UpdatedAt  time.Time
}

type Attachment struct {
	ID          uuid.UUID
	QuestionID  uuid.UUID
	FilePath    string
	IsPersisted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Answer struct {
	ID               uuid.UUID
	QuestionID       uuid.UUID
	Question         Question
	UserID           uuid.UUID
	AnswerText       string
	SelectedOptionID *uuid.UUID
	SelectedOption   *Option
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

/// validations

func validateTitle(title string) error {
	if title == "" {
		return ErrTitleEmpty
	}
	if len(title) > 255 {
		return ErrExceedsMinCharacter
	}

	return nil
}
