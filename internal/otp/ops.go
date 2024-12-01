package otp

import (
	"context"

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
func (o *Ops) Create(ctx context.Context, otp *OTP) (*OTP, error) {
	createdOTP, err := o.repo.Create(ctx, otp)
	if err != nil {
		return nil, ErrNotCreate
	}
	return createdOTP, nil
}
func (o *Ops) GetByUserIdAndCode(ctx context.Context, userId uuid.UUID, code string) (*OTP, error) {
	otp, err := o.repo.GetByUserIdAndCode(ctx, userId, code)
	if err != nil {
		return nil, ErrNotRetreived
	}
	return otp, nil
}
