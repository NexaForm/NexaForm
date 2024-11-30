package storage

import (
	"NexaForm/internal/role"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"NexaForm/pkg/utils"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) role.Repo {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) Create(ctx context.Context, role *role.Role) (*role.Role, error) {
	newRole := mappers.RoleDomainToEntity(role)
	err := r.db.WithContext(ctx).Create(&newRole).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, utils.DbErrDuplicateKey
		}
		return nil, err
	}
	cratedRole := mappers.RoleEntityToDomain(newRole)
	return cratedRole, nil
}
func (r *roleRepo) GetByID(ctx context.Context, id uuid.UUID) (*role.Role, error) {
	var roleEntity entities.Role
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&roleEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// Map entity to domain
	return mappers.RoleEntityToDomain(&roleEntity), nil
}
func (r *roleRepo) GetRoleByName(ctx context.Context, name string) (*role.Role, error) {
	var roleEntity entities.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&roleEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// Map entity to domain
	return mappers.RoleEntityToDomain(&roleEntity), nil
}
