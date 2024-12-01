package role

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRoleAlreadyExists = errors.New("email already exists")
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
