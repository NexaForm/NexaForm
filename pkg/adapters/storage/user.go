package storage

import (
	"NexaForm/internal/user"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"NexaForm/pkg/utils"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	newUser := mappers.UserDomainToEntity(user)
	err := r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, utils.DbErrDuplicateKey
		}
		return nil, err
	}
	createdUser := mappers.UserEntityToDomain(newUser)
	return createdUser, nil
}

func (r *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var u entities.User

	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&u), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}

func (r *userRepo) ActivateUser(ctx context.Context, email string) (*user.User, error) {
	var user entities.User

	query := `
		UPDATE users
		SET is_email_verified = TRUE, updated_at = NOW()
		WHERE email = $1
		RETURNING *;
	`

	// Execute the query and scan the result into the user object
	if err := r.db.WithContext(ctx).Raw(query, email).Scan(&user).Error; err != nil {
		return nil, err
	}

	verifiedUser := mappers.UserEntityToDomain(&user)
	return verifiedUser, nil
}


func (r *userRepo)	GetUserByIDWithNumberOfHisSurveys(ctx context.Context, id uuid.UUID) (*user.User, int,error){
	var u entities.User

	err := r.db.WithContext(ctx).Model(&entities.User{}).Preload("Surveys").Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,0, nil
		}
		return nil,0, err
	}

	return mappers.UserFullEntityToDomain(&u), len(u.Surveys), nil
}

func (r *userRepo) GetAllVerifiedUsers(ctx context.Context, limit, offset uint) ([]user.User, uint, error) {
	var total int64
	query := r.db.WithContext(ctx).Model(&entities.User{}).Where("is_email_verified = ?", true)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if offset > 0 {
		query = query.Offset(int(offset))
	}

	if limit > 0 {
		query = query.Limit(int(limit))
	}

	var users []entities.User

	if err := query.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	domainUsers := mappers.BatchUserEntityToDomain(users)
	return domainUsers, uint(total), nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	var existingUser entities.User
	if err := r.db.WithContext(ctx).First(&existingUser, "id = ?", user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	if user.FullName != "" {
		existingUser.FullName = user.FullName

	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.NationalID != "" {
		existingUser.NationalID = user.NationalID

	}

	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return nil, err
	}

	updatedUser := mappers.UserEntityToDomain(&existingUser)
	return updatedUser, nil
}
