package handlers

import (
	presenter "NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/survey"
	"NexaForm/pkg/jwt"
	"NexaForm/service"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddSurveyHandler(surveyService *service.SurveyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.AddSurveyRequest
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}
		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		survey := presenter.AddSurveyRequestToSurveyDomain(&req)
		createdSurvey, err := surveyService.CreateSurvey(c.Context(), survey)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}
		return presenter.Created(c, "Survey created successfully", createdSurvey)
	}
}

// GetPresignedURLsHandler handles the generation of presigned URLs for survey questions' attachments
func GetPresignedURLsHandler(surveyService *service.SurveyService, fileService *service.FileService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse survey ID from query parameter
		surveyIDStr := c.Query("survey_id")
		if surveyIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey_id is required"))
		}

		// Convert survey ID to UUID
		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey_id format"))
		}

		// Fetch questions for the survey
		questions, err := surveyService.GetQuestionsBySurveyID(c.Context(), surveyID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}
		if len(questions) == 0 {
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no questions found for the survey"))
		}

		// Parse file metadata from the request body
		var fileRequests = presenter.FileRequests
		if err := c.BodyParser(&fileRequests); err != nil {
			return presenter.BadRequest(c, err)
		}

		// Prepare object paths and validate content types
		attachments := make([]survey.Attachment, len(fileRequests))
		objectPaths := make([]string, len(fileRequests))
		contentTypes := make([]string, len(fileRequests))
		for i, req := range fileRequests {
			if req.FileName == "" {
				return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "file_name is required"))
			}
			attachmentID := uuid.New()
			attachments[i] = survey.Attachment{
				ID:          attachmentID,
				QuestionID:  req.QuestionID,
				FilePath:    "attachments/" + req.QuestionID.String() + "/" + req.FileName,
				IsPersisted: false, // Initially not persisted
			}
			objectPaths[i] = attachments[i].FilePath
			contentTypes[i] = req.ContentType
		}

		// Generate presigned URLs
		presignedURLs, err := fileService.GeneratePresignedUploadURLs(c.Context(), objectPaths, contentTypes)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		// Create attachments in the database
		err = surveyService.CreateAttachments(c.Context(), attachments...)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		// Return the presigned URLs
		return presenter.OK(c, "Presigned URLs generated successfully", presignedURLs)
	}
}

// GetPresignedDownloadURLsHandler handles the generation of presigned URLs for downloading attachments
func GetPresignedDownloadURLsHandler(fileService *service.FileService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse survey ID from query parameter
		surveyIDStr := c.Query("survey_id")
		if surveyIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey_id is required"))
		}

		// Convert survey ID to UUID
		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey_id format"))
		}

		// Generate presigned URLs for attachments
		presignedURLs, err := fileService.GeneratePresignedDownloadURLs(c.Context(), surveyID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		// Return the presigned URLs
		return presenter.OK(c, "Presigned download URLs generated successfully", presignedURLs)
	}
}
func CreateAnswerHandler(surveyService *service.SurveyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateAnswerRequest
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		// Validate request
		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		// Extract user information from JWT claims
		claims := c.Locals(jwt.UserClaimKey)
		userClaims, ok := claims.(*jwt.UserClaims)
		if !ok {
			return presenter.Unauthorized(c, errors.New("invalid user claims"))
		}

		// Prepare answer
		answer := survey.Answer{
			QuestionID:       req.QuestionID,
			UserID:           userClaims.UserID,
			AnswerText:       req.AnswerText,
			SelectedOptionID: req.SelectedOptionID,
		}

		// Call the service to handle the answer and determine the next question
		nextQuestion, err := surveyService.CreateAnswer(c.Context(), answer)
		if err != nil {
			if errors.Is(err, service.ErrPreviousQuestionUnanswered) {
				return presenter.Unauthorized(c, err)
			}
			if strings.Contains(err.Error(), "already been answered") {
				return presenter.Conflict(c, err)
			}
			return presenter.InternalServerError(c, err)
		}
		if nextQuestion == nil {
			return presenter.OK(c, "Survey completed", nil)
		}
		// Return the next question
		return presenter.OK(c, "Next question retrieved successfully", nextQuestion)
	}
}
