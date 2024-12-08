package mappers

import (
	"NexaForm/internal/rbac"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/fp"
)

// survey role mappers
func SurveyRoleEntityToDomain(entity *entities.SurveyRole) *rbac.SurveyRole {
	return &rbac.SurveyRole{
		ID:           entity.ID,
		SurveyID:     entity.SurveyID,
		Name:         entity.Name,
		Permissions:  BatchSurveyPermissionEntityToDomain(entity.Permissions),
		Participants: BatchSurveyParticipantEntityToDomain(entity.Participants),
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
func surveyRoleEntityToDomain(entity entities.SurveyRole) rbac.SurveyRole {
	return rbac.SurveyRole{
		ID:           entity.ID,
		SurveyID:     entity.SurveyID,
		Name:         entity.Name,
		Permissions:  BatchSurveyPermissionEntityToDomain(entity.Permissions),
		Participants: BatchSurveyParticipantEntityToDomain(entity.Participants),
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
func BatchSurveyRoleEntityToDomain(entities []entities.SurveyRole) []rbac.SurveyRole {
	return fp.Map(entities, surveyRoleEntityToDomain)
}
func SurveyRoleDomainToEntity(domain *rbac.SurveyRole) *entities.SurveyRole {
	return &entities.SurveyRole{
		ID:           domain.ID,
		SurveyID:     domain.SurveyID,
		Name:         domain.Name,
		Permissions:  BatchSurveyPermissionDomainToEntity(domain.Permissions),
		Participants: BatchSurveyParticipantDomainToEntity(domain.Participants),
	}
}
func surveyRoleDomainToEntity(domain rbac.SurveyRole) entities.SurveyRole {
	return entities.SurveyRole{
		ID:           domain.ID,
		SurveyID:     domain.SurveyID,
		Name:         domain.Name,
		Permissions:  BatchSurveyPermissionDomainToEntity(domain.Permissions),
		Participants: BatchSurveyParticipantDomainToEntity(domain.Participants),
	}
}
func BatchSurveyRoleDomainToEntity(domains []rbac.SurveyRole) []entities.SurveyRole {
	return fp.Map(domains, surveyRoleDomainToEntity)
}

// survey permissions mappers

func SurveyPermissionEntityToDomain(entity *entities.SurveyPermission) *rbac.SurveyPermission {
	return &rbac.SurveyPermission{
		ID:                   entity.ID,
		SurveyID:             entity.SurveyID,
		Name:                 entity.Name,
		SurveyRoles:          BatchSurveyRoleEntityToDomain(entity.SurveyRoles),
		Description:          entity.Description,
		CanWatchSurvey:       entity.CanWatchSurvey,
		CanWatchExposedVotes: entity.CanWatchExposedVotes,
		CanVote:              entity.CanVote,
		CanEditSurvey:        entity.CanEditSurvey,
		CanAssignRole:        entity.CanAssignRole,
		CanAccessReports:     entity.CanAccessReports,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}
}
func surveyPermissionEntityToDomain(entity entities.SurveyPermission) rbac.SurveyPermission {
	return rbac.SurveyPermission{
		ID:                   entity.ID,
		SurveyID:             entity.SurveyID,
		Name:                 entity.Name,
		SurveyRoles:          BatchSurveyRoleEntityToDomain(entity.SurveyRoles),
		Description:          entity.Description,
		CanWatchSurvey:       entity.CanWatchSurvey,
		CanWatchExposedVotes: entity.CanWatchExposedVotes,
		CanVote:              entity.CanVote,
		CanEditSurvey:        entity.CanEditSurvey,
		CanAssignRole:        entity.CanAssignRole,
		CanAccessReports:     entity.CanAccessReports,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}
}
func BatchSurveyPermissionEntityToDomain(entities []entities.SurveyPermission) []rbac.SurveyPermission {
	return fp.Map(entities, surveyPermissionEntityToDomain)
}
func SurveyPermissionDomainToEntity(domain *rbac.SurveyPermission) *entities.SurveyPermission {
	return &entities.SurveyPermission{
		ID:                   domain.ID,
		SurveyID:             domain.SurveyID,
		Name:                 domain.Name,
		SurveyRoles:          BatchSurveyRoleDomainToEntity(domain.SurveyRoles),
		Description:          domain.Description,
		CanWatchSurvey:       domain.CanWatchSurvey,
		CanWatchExposedVotes: domain.CanWatchExposedVotes,
		CanVote:              domain.CanVote,
		CanEditSurvey:        domain.CanEditSurvey,
		CanAssignRole:        domain.CanAssignRole,
		CanAccessReports:     domain.CanAccessReports,
	}
}
func surveyPermissionDomainToEntity(domain rbac.SurveyPermission) entities.SurveyPermission {
	return entities.SurveyPermission{
		ID:                   domain.ID,
		SurveyID:             domain.SurveyID,
		Name:                 domain.Name,
		SurveyRoles:          BatchSurveyRoleDomainToEntity(domain.SurveyRoles),
		Description:          domain.Description,
		CanWatchSurvey:       domain.CanWatchSurvey,
		CanWatchExposedVotes: domain.CanWatchExposedVotes,
		CanVote:              domain.CanVote,
		CanEditSurvey:        domain.CanEditSurvey,
		CanAssignRole:        domain.CanAssignRole,
		CanAccessReports:     domain.CanAccessReports,
	}
}
func BatchSurveyPermissionDomainToEntity(domains []rbac.SurveyPermission) []entities.SurveyPermission {
	return fp.Map(domains, surveyPermissionDomainToEntity)
}

// survey participants mappers

func SurveyParticipantEntityToDomain(entity *entities.SurveyParticipant) *rbac.SurveyParticipant {
	return &rbac.SurveyParticipant{
		ID:           entity.ID,
		SurveyID:     entity.SurveyID,
		UserID:       entity.UserID,
		SurveyRoleID: entity.SurveyRoleID,
		IsExposed:    entity.IsExposed,
		RoleExpire:   entity.RoleExpire,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
func surveyParticipantEntityToDomain(entity entities.SurveyParticipant) rbac.SurveyParticipant {
	return rbac.SurveyParticipant{
		ID:           entity.ID,
		SurveyID:     entity.SurveyID,
		UserID:       entity.UserID,
		SurveyRoleID: entity.SurveyRoleID,
		IsExposed:    entity.IsExposed,
		RoleExpire:   entity.RoleExpire,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
func BatchSurveyParticipantEntityToDomain(entities []entities.SurveyParticipant) []rbac.SurveyParticipant {
	return fp.Map(entities, surveyParticipantEntityToDomain)
}
func SurveyParticipantDomainToEntity(domain *rbac.SurveyParticipant) *entities.SurveyParticipant {
	return &entities.SurveyParticipant{
		ID:           domain.ID,
		SurveyID:     domain.SurveyID,
		UserID:       domain.UserID,
		SurveyRoleID: domain.SurveyRoleID,
		IsExposed:    domain.IsExposed,
		RoleExpire:   domain.RoleExpire,
	}
}
func surveyParticipantDomainToEntity(domain rbac.SurveyParticipant) entities.SurveyParticipant {
	return entities.SurveyParticipant{
		ID:           domain.ID,
		SurveyID:     domain.SurveyID,
		UserID:       domain.UserID,
		SurveyRoleID: domain.SurveyRoleID,
		IsExposed:    domain.IsExposed,
		RoleExpire:   domain.RoleExpire,
	}
}
func BatchSurveyParticipantDomainToEntity(domains []rbac.SurveyParticipant) []entities.SurveyParticipant {
	return fp.Map(domains, surveyParticipantDomainToEntity)
}
