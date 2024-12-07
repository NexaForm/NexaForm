package survey

import (
	"NexaForm/internal/user"
	"context"
	"time"

	"github.com/google/uuid"
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
}
type Survey struct {
	ID                 uuid.UUID
	OwnerID            uuid.UUID
	Title              string
	Description        string
	StartTime          time.Time
	EndTime            time.Time
	Visibility         VisibilityType
	AllowedMinAge      int
	AllowedMaxAge      int
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
	ID            uuid.UUID
	SurveyID      uuid.UUID
	Description   string
	Type          QuestionType
	Order         int
	IsConditional bool
	Options       []Option
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Option struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Text       string
	IsCorrect  *bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
