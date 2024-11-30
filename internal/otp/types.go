package otp

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotCreate    = errors.New("otp not created")
	ErrNotRetreived = errors.New("otp not retreived")
)

type Repo interface {
	Create(ctx context.Context, otp *OTP) (*OTP, error)
	GetByUserIdAndCode(ctx context.Context, userId uuid.UUID, code string) (*OTP, error)
}

type OTP struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	OTPCode   string
	OTPExpiry time.Time
}
