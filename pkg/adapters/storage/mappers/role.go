package mappers

import (
	"NexaForm/internal/role"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/fp"
)

func roleEntityToDomain(entity entities.Role) role.Role {
	return role.Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func BatchRoleEntityToDomain(entities []entities.Role) []role.Role {
	return fp.Map(entities, roleEntityToDomain)
}

// RoleDomainToEntity maps domain model to database entity
func RoleDomainToEntity(domain *role.Role) *entities.Role {
	return &entities.Role{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

// RoleEntityToDomain maps database entity to domain model
func RoleEntityToDomain(entity *entities.Role) *role.Role {
	return &role.Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
