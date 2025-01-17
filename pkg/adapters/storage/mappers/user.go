package mappers

import (
	"NexaForm/internal/role"
	"NexaForm/internal/user"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/fp"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:           entity.ID,
		FullName:     entity.FullName,
		Email:        entity.Email,
		EmailIsValid: *entity.IsEmailVerified,
		Password:     entity.Password,
		Role:         role.Role{ID: entity.RoleID, Name: entity.Role.Name},
	}
}
func userEntityToDomain(entity entities.User) user.User {
	return user.User{
		ID:           entity.ID,
		FullName:     entity.FullName,
		Email:        entity.Email,
		EmailIsValid: *entity.IsEmailVerified,
		Password:     entity.Password,
		Role:         role.Role{ID: entity.RoleID, Name: entity.Role.Name},
	}
}

func BatchUserEntityToDomain(entities []entities.User) []user.User {
	return fp.Map(entities, userEntityToDomain)
}

func UserDomainToEntity(domainUser *user.User) *entities.User {
	return &entities.User{
		RoleID:          domainUser.Role.ID,
		Role:            entities.Role{ID: domainUser.Role.ID, Name: domainUser.Role.Name},
		FullName:        domainUser.FullName,
		Email:           domainUser.Email,
		IsEmailVerified: &domainUser.EmailIsValid,
		NationalID:      domainUser.NationalID,
		Password:        domainUser.Password,
	}
}
