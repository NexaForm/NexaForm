package handlers

import (
	"NexaForm/api/http/handlers/presenter"
	"NexaForm/internal/rbac"
	"NexaForm/service"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateSurveyRoleHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateRoleReq
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		roles := make([]rbac.SurveyRole, len(req.Role))
		for i, r := range req.Role {
			roles[i] = rbac.SurveyRole{
				SurveyID: r.SurveyID,
				Name:     r.Name,
			}
		}

		createdRoles, err := rbacService.CreateSurveyRoles(c.Context(), roles)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		return presenter.Created(c, "Survey roles created successfully", presenter.BatchSurveyRoleToResponse(createdRoles))
	}
}
func GetSurveyRoleHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleIDStr := c.Params("id")
		if roleIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "role ID is required"))
		}

		roleID, err := uuid.Parse(roleIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid role ID format"))
		}

		role, err := rbacService.GetSurveyRole(c.Context(), roleID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		if role == nil {
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "survey role not found"))
		}

		return presenter.OK(c, "Survey role retrieved successfully", presenter.SurveyRoleToResponse(role))
	}
}
func GetSurveyRolesBySurveyIDHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		surveyIDStr := c.Params("survey_id")
		if surveyIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey ID is required"))
		}

		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey ID format"))
		}

		roles, err := rbacService.GetSurveyRolesBySurveyID(c.Context(), surveyID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		if len(roles) == 0 {
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no roles found for the survey"))
		}

		return presenter.OK(c, "Survey roles retrieved successfully", presenter.BatchSurveyRoleToResponse(roles))
	}
}
func CreateSurveyPermissionHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateSurveyPermissionsRequest
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		permissions := make([]rbac.SurveyPermission, len(req))
		for i, p := range req {
			permissions[i] = rbac.SurveyPermission{
				SurveyID:             p.SurveyID,
				Name:                 p.Name,
				Description:          sql.NullString{String: p.Description, Valid: p.Description != ""},
				CanWatchSurvey:       p.CanWatchSurvey,
				CanWatchExposedVotes: p.CanWatchExposedVotes,
				CanVote:              p.CanVote,
				CanEditSurvey:        p.CanEditSurvey,
				CanAssignRole:        p.CanAssignRole,
				CanAccessReports:     p.CanAccessReports,
			}
		}

		createdPermissions, err := rbacService.CreateSurveyPermissions(c.Context(), permissions)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		return presenter.Created(c, "Survey permissions created successfully", presenter.BatchSurveyPermissionToResponse(createdPermissions))
	}
}
func GetSurveyPermissionsBySurveyIDHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		surveyIDStr := c.Params("survey_id")
		if surveyIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey ID is required"))
		}

		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey ID format"))
		}

		permissions, err := rbacService.GetSurveyPermissionsBySurveyID(c.Context(), surveyID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		if len(permissions) == 0 {
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no permissions found for the survey"))
		}

		return presenter.OK(c, "Survey permissions retrieved successfully", presenter.BatchSurveyPermissionToResponse(permissions))
	}
}
func CreateSurveyParticipantHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateSurveyParticipantsRequest
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		participants := make([]rbac.SurveyParticipant, len(req))
		for i, p := range req {
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
			return presenter.InternalServerError(c, err)
		}

		return presenter.Created(c, "Survey participants created successfully", presenter.BatchSurveyParticipantToResponse(createdParticipants))
	}
}
func GetSurveyParticipantsBySurveyIDHandler(rbacService *service.RBACService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		surveyIDStr := c.Params("survey_id")
		if surveyIDStr == "" {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "survey ID is required"))
		}

		surveyID, err := uuid.Parse(surveyIDStr)
		if err != nil {
			return presenter.BadRequest(c, fiber.NewError(fiber.StatusBadRequest, "invalid survey ID format"))
		}

		participants, err := rbacService.GetSurveyParticipantsBySurveyID(c.Context(), surveyID)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		if len(participants) == 0 {
			return presenter.NotFound(c, fiber.NewError(fiber.StatusNotFound, "no participants found for the survey"))
		}

		return presenter.OK(c, "Survey participants retrieved successfully", presenter.BatchSurveyParticipantToResponse(participants))
	}
}
