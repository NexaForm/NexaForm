package service

import (
	"NexaForm/internal/survey"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var ErrPreviousQuestionUnanswered = errors.New("you must answer the previous question before proceeding")

type SurveyService struct {
	surveyOps *survey.Ops
}

// NewSurveyService initializes SurveyService and sets up Redis client
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

func (ss *SurveyService) GetQuestionsBySurveyID(ctx context.Context, surveyId uuid.UUID) ([]survey.Question, error) {
	return ss.surveyOps.GetQuestionsBySurveyID(ctx, surveyId)
}

func (ss *SurveyService) CreateAttachments(ctx context.Context, attachments ...survey.Attachment) error {
	return ss.surveyOps.CreateAttachments(ctx, attachments...)
}

func (ss *SurveyService) UpdateAttachments(ctx context.Context, attachments ...survey.Attachment) error {
	return ss.surveyOps.UpdateAttachments(ctx, attachments...)
}
func (ss *SurveyService) CreateAnswer(ctx context.Context, answer survey.Answer) (*survey.Question, error) {
	// Check if the question has already been answered
	existingAnswer, err := ss.surveyOps.CheckAnswerExists(ctx, answer.QuestionID, answer.UserID)
	if err != nil {
		return nil, err
	}
	if existingAnswer != nil {
		return nil, fmt.Errorf("question %s has already been answered", answer.QuestionID.String())
	}

	// Fetch the survey by question ID along with its questions
	s, err := ss.surveyOps.GetSurveyByQuestionID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}

	// Find the current question and check if it has a target question
	var currentQuestion *survey.Question
	for _, question := range s.Questions {
		if question.ID == answer.QuestionID {
			currentQuestion = &question
			break
		}
	}
	if currentQuestion == nil {
		return nil, fmt.Errorf("question %s not found in survey %s", answer.QuestionID.String(), s.ID.String())
	}

	// Check if the survey is ordered and validate previous questions
	if s.IsOrdered {
		answeredQuestions, err := ss.surveyOps.GetAnsweredQuestionsByUser(ctx, s.ID, answer.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch answered questions: %w", err)
		}

		answeredMap := make(map[uuid.UUID]bool)
		for _, q := range answeredQuestions {
			answeredMap[q.ID] = true
		}

		// Ensure all questions with a smaller `Order` are answered
		for _, question := range s.Questions {
			if question.Order < currentQuestion.Order && !answeredMap[question.ID] {
				return nil, ErrPreviousQuestionUnanswered
			}
		}
	}

	// Save the answer to the database
	_, err = ss.surveyOps.CreateAnswer(ctx, answer)
	if err != nil {
		return nil, fmt.Errorf("failed to save answer: %w", err)
	}

	// Check if the current question triggers a target question
	if currentQuestion.TargetQuestionID != nil {
		for _, question := range s.Questions {
			if question.ID == *currentQuestion.TargetQuestionID {
				return &question, nil // Return the target question
			}
		}
		return nil, fmt.Errorf("target question %s not found", currentQuestion.TargetQuestionID.String())
	}

	// Handle ordered and unordered survey logic
	nextQuestion, err := ss.getNextQuestion(ctx, s, answer.UserID)
	if err != nil {
		return nil, err
	}
	// If no more questions, return a special signal (nil, nil)
	if nextQuestion == nil {
		return nil, nil // Survey is completed
	}

	return nextQuestion, nil
}

func (ss *SurveyService) getNextQuestion(ctx context.Context, s *survey.Survey, userID uuid.UUID) (*survey.Question, error) {
	// Fetch answered question IDs
	answeredQuestions, err := ss.surveyOps.GetAnsweredQuestionsByUser(ctx, s.ID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answered questions: %w", err)
	}

	answeredMap := make(map[uuid.UUID]bool)
	for _, q := range answeredQuestions {
		answeredMap[q.ID] = true
	}

	// Separate answered and unanswered questions
	var unansweredQuestions []survey.Question
	for _, question := range s.Questions {
		if !answeredMap[question.ID] && !question.IsConditional {
			unansweredQuestions = append(unansweredQuestions, question)
		}
	}

	if s.IsOrdered {
		// Handle ordered surveys: find the next question with the smallest order
		var nextQuestion *survey.Question
		currentOrder := int(^uint(0) >> 1) // Max int value

		for _, question := range unansweredQuestions {
			if question.Order < currentOrder {
				nextQuestion = &question
				currentOrder = question.Order
			}
		}

		return nextQuestion, nil
	} else {
		if len(unansweredQuestions) > 0 {
			randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
			randomIndex := randomGen.Intn(len(unansweredQuestions))
			return &unansweredQuestions[randomIndex], nil
		}
	}

	return nil, nil
}
