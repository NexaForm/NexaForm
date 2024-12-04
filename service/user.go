package service

import (
	"NexaForm/internal/user"
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

func (s *UserService) GetAllVerifiedUsers(ctx context.Context, userID uuid.UUID, page, pageSize uint) ([]user.User, uint, error) {
	_, err := s.userOps.GetUserByID(ctx, userID)
	if err != nil {
		return nil, 0, err
	}

	return s.userOps.GetAllVerifiedUsers(ctx, page, pageSize)
}
