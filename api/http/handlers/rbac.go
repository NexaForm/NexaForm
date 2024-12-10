package handlers

import (
	"NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/rbac"
	"NexaForm/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateSurveyRoleHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("CreateSurveyRoleHandler")
		var req presenter.CreateRoleReq
		if err := c.BodyParser(&req); err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}

		err := BodyValidator(req)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}

		roles := make([]rbac.SurveyRole, len(req.Role))
		for i, r := range req.Role {
			roles[i] = rbac.SurveyRole{
				SurveyID:             r.SurveyID,
				Name:                 r.Name,
				CanWatchSurvey:       r.CanWatchSurvey,
				CanWatchExposedVotes: r.CanWatchExposedVotes,
				CanEditSurvey:        r.CanEditSurvey,
				CanVote:              r.CanVote,
				CanAssignRole:        r.CanAssignRole,
				CanAccessReports:     r.CanAccessReports,
			}
		}

		createdRoles, err := rbacService.CreateSurveyRoles(c.Context(), roles)
		if err != nil {
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}
		logger.Info("Survey roles created successfully")
		return presenter.Created(c, "Survey roles created successfully", presenter.BatchSurveyRoleToResponse(createdRoles))
	}
}
func GetSurveyRoleHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("GetSurveyRoleHandler")
		roleIDStr := c.Params("id")
		if roleIDStr == "" {
			logger.Error("role ID is required")
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "role ID is required"))
		}

		roleID, err := uuid.Parse(roleIDStr)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid role ID format"))
		}

		role, err := rbacService.GetSurveyRole(c.Context(), roleID)
		if err != nil {
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}

		if role == nil {
			logger.Error("survey role not found")
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "survey role not found"))
		}
		logger.Info("Survey role retrieved successfully")
		return presenter.OK(c, "Survey role retrieved successfully", presenter.SurveyRoleToResponse(role))
	}
}
func GetSurveyRolesBySurveyIDHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("GetSurveyRolesBySurveyIDHandler")
		surveyIDStr := c.Params("survey_id")
		if surveyIDStr == "" {
			logger.Error("survey ID is required")
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey ID is required"))
		}

		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey ID format"))
		}

		roles, err := rbacService.GetSurveyRolesBySurveyID(c.Context(), surveyID)
		if err != nil {
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}

		if len(roles) == 0 {
			logger.Error("no roles found for the survey")
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no roles found for the survey"))
		}
		logger.Info("Survey roles retrieved successfully")
		return presenter.OK(c, "Survey roles retrieved successfully", presenter.BatchSurveyRoleToResponse(roles))
	}
}

func CreateSurveyParticipantHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("CreateSurveyParticipantHandler")
		var req presenter.CreateParticipantReq
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		err := BodyValidator(req)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, err)
		}

		participants := make([]rbac.SurveyParticipant, len(req.Participant))
		for i, p := range req.Participant {
			participants[i] = rbac.SurveyParticipant{
				SurveyID:     p.SurveyID,
				UserID:       p.UserID,
				SurveyRoleID: p.SurveyRoleID,
				IsExposed:    p.IsExposed,
				RoleExpire:   p.RoleExpire,
			}
		}

		createdParticipants, err := rbacService.CreateSurveyParticipants(c.Context(), participants)
		if err != nil {
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}
		logger.Info("Survey participants created successfully")
		return presenter.Created(c, "Survey participants created successfully", presenter.BatchSurveyParticipantToResponse(createdParticipants))
	}
}
func GetSurveyParticipantsBySurveyIDHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := GetLogger(c)
		defer logger.Sync()
		logger.Info("GetSurveyParticipantsBySurveyIDHandler")
		surveyIDStr := c.Params("survey_id")
		if surveyIDStr == "" {
			logger.Error("survey ID is required")
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey ID is required"))
		}

		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			logger.Error(err.Error())
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey ID format"))
		}

		participants, err := rbacService.GetSurveyParticipantsBySurveyID(c.UserContext(), surveyID)
		if err != nil {
			logger.Error(err.Error())
			return presenter.InternalServerError(c, err)
		}

		if len(participants) == 0 {
			logger.Error("no participants found for the survey")
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no participants found for the survey"))
		}
		logger.Info("Survey participants retrieved successfully")
		return presenter.OK(c, "Survey participants retrieved successfully", presenter.BatchSurveyParticipantToResponse(participants))
	}
}
