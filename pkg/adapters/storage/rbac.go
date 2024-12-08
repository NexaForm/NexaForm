package storage

import (
	"NexaForm/internal/rbac"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rbacRepo struct {
	db *gorm.DB
}

// NewRBACRepo creates a new instance of the RBAC repository
func NewRBACRepo(db *gorm.DB) rbac.Repo {
	return &rbacRepo{
		db: db,
	}
}

// CreateSurveyRoles inserts multiple survey roles into the database
func (r *rbacRepo) CreateSurveyRoles(ctx context.Context, surveyRoles []rbac.SurveyRole) ([]rbac.SurveyRole, error) {
	entities := mappers.BatchSurveyRoleDomainToEntity(surveyRoles)
	err := r.db.WithContext(ctx).Create(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyRoleEntityToDomain(entities), nil
}

// CreateSurveyPermissions inserts multiple survey permissions into the database
func (r *rbacRepo) CreateSurveyPermissions(ctx context.Context, surveyPermissions []rbac.SurveyPermission) ([]rbac.SurveyPermission, error) {
	entities := mappers.BatchSurveyPermissionDomainToEntity(surveyPermissions)
	err := r.db.WithContext(ctx).Create(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyPermissionEntityToDomain(entities), nil
}

// CreateSurveyParticipants inserts multiple survey participants into the database
func (r *rbacRepo) CreateSurveyParticipants(ctx context.Context, surveyParticipants []rbac.SurveyParticipant) ([]rbac.SurveyParticipant, error) {
	entities := mappers.BatchSurveyParticipantDomainToEntity(surveyParticipants)
	err := r.db.WithContext(ctx).Create(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyParticipantEntityToDomain(entities), nil
}

// GetSurveyRole retrieves a survey role by its ID
func (r *rbacRepo) GetSurveyRole(ctx context.Context, id uuid.UUID) (*rbac.SurveyRole, error) {
	var entity entities.SurveyRole
	err := r.db.WithContext(ctx).Preload("Permissions").Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mappers.SurveyRoleEntityToDomain(&entity), nil
}

// GetSurveyPermission retrieves a survey permission by its ID
func (r *rbacRepo) GetSurveyPermission(ctx context.Context, id uuid.UUID) (*rbac.SurveyPermission, error) {
	var entity entities.SurveyPermission
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mappers.SurveyPermissionEntityToDomain(&entity), nil
}

// GetSurveyParticipant retrieves a survey participant by its ID
func (r *rbacRepo) GetSurveyParticipant(ctx context.Context, id uuid.UUID) (*rbac.SurveyParticipant, error) {
	var entity entities.SurveyParticipant
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mappers.SurveyParticipantEntityToDomain(&entity), nil
}

// GetSurveyRolesBySurveyID retrieves all survey roles for a specific survey
func (r *rbacRepo) GetSurveyRolesBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyRole, error) {
	var entities []entities.SurveyRole
	err := r.db.WithContext(ctx).Preload("Permissions").Where("survey_id = ?", surveyID).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyRoleEntityToDomain(entities), nil
}

// GetSurveyPermissionsBySurveyID retrieves all survey permissions for a specific survey
func (r *rbacRepo) GetSurveyPermissionsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyPermission, error) {
	var entities []entities.SurveyPermission
	err := r.db.WithContext(ctx).Where("survey_id = ?", surveyID).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyPermissionEntityToDomain(entities), nil
}

// GetSurveyParticipantsBySurveyID retrieves all survey participants for a specific survey
func (r *rbacRepo) GetSurveyParticipantsBySurveyID(ctx context.Context, surveyID uuid.UUID) ([]rbac.SurveyParticipant, error) {
	var entities []entities.SurveyParticipant
	err := r.db.WithContext(ctx).Where("survey_id = ?", surveyID).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return mappers.BatchSurveyParticipantEntityToDomain(entities), nil
}
