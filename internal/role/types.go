package role

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidEmail          = errors.New("invalid email format")
	ErrInvalidPassword       = errors.New("invalid password format")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrInvalidAuthentication = errors.New("email and password doesn't match")
)

type Repo interface {
	Create(ctx context.Context, role *Role) (*Role, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Role, error)
	GetRoleByName(ctx context.Context, name string) (*Role, error)
}

type Role struct {
	ID   uuid.UUID
	Name string
}
