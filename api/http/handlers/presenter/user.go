package presenter

import (
	"NexaForm/internal/user"
	"NexaForm/pkg/fp"

	"github.com/google/uuid"
)

type UserGet struct {
	ID       uuid.UUID `json:"user_id"`
	FullName *string   `json:"full_name"`
	Email    string    `json:"email"`
}

func UserToUserGet(user user.User) UserGet {
	return UserGet{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

func BatchUsersToUserGet(users []user.User) []UserGet {
	return fp.Map(users, UserToUserGet)
}
