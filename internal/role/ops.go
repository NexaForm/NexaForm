package role

import (
	"NexaForm/pkg/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, role *Role) (*Role, error) {
	createdRole, err := o.repo.Create(ctx, role)
	if err != nil {
		if errors.Is(err, utils.DbErrDuplicateKey) {
			return nil, ErrRoleAlreadyExists
		}
		return nil, err
	}
	return createdRole, nil
}

func (o *Ops) GetRoleByID(ctx context.Context, id uuid.UUID) (*Role, error) {
	return o.repo.GetByID(ctx, id)
}
func (o *Ops) GetRoleByName(ctx context.Context, name string) (*Role, error) {
	return o.repo.GetRoleByName(ctx, name)
}
