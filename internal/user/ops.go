package user

import (
	"NexaForm/internal/role"
	"NexaForm/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Ops struct {
	repo     Repo
	roleRepo role.Repo
}

func NewOps(repo Repo, roleRepo role.Repo) *Ops {
	return &Ops{
		repo:     repo,
		roleRepo: roleRepo,
	}
}

func (o *Ops) Create(ctx context.Context, user *User) (*User, error) {

	err := validateUserRegistration(user)
	if err != nil {
		return nil, err
	}
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)
	// Fetch default role if not provided
	if user.Role.ID == uuid.Nil {
		defaultRole, err := o.roleRepo.GetRoleByName(ctx, "User")
		if err != nil {
			log.Printf("Error fetching default role: %v", err)
			return nil, fmt.Errorf("failed to fetch default role: %w", err)
		}
		log.Printf("Default role fetched: %v", defaultRole)
		user.Role.ID = defaultRole.ID
	}
	log.Printf("User Role ID: %v", user.Role.ID)
	// lowercase email
	user.Email = LowerCaseEmail(user.Email)
	createdUser, err := o.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, utils.DbErrDuplicateKey) {
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}
	return createdUser, nil
}

func (o *Ops) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {

	user, err := o.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, err
}

func (o *Ops) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	email = LowerCaseEmail(email)
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return nil, ErrInvalidAuthentication
	}

	return user, nil
}

func (o *Ops) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	email = LowerCaseEmail(email)
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (o *Ops) ActivateUser(ctx context.Context, email string) (*User, error) {
	user, err := o.repo.ActivateUser(ctx, email)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (o *Ops) GetAllVerifiedUsers(ctx context.Context, page, pageSize uint) ([]User, uint, error) {
	limit := pageSize
	offset := (page - 1) * pageSize

	return o.repo.GetAllVerifiedUsers(ctx, limit, offset)
}

func validateUserRegistration(user *User) error {
	err := ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	if err := ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	return nil
}

func (o *Ops) UpdateUser(ctx context.Context, user *User) (*User, error) {
	if err := ValidateEmail(user.Email); err != nil {
		return nil, err
	}

	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.SetPassword(hashedPassword)
	}

	user.Email = LowerCaseEmail(user.Email)

	updatedUser, err := o.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
