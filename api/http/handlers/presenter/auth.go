package presenter

import (
	"NexaForm/internal/role"
	"NexaForm/internal/user"
)

type UserRegisterReq struct {
	FullName   string `json:"full_name" validate:"required" example:"yourname"`
	Email      string `json:"email" validate:"required" example:"abc@gmail.com"`
	NationalID string `json:"national_id"`
	Password   string `json:"password" validate:"required" example:"Abc@123"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required" example:"valid_email@folan.com"`
	Password string `json:"password" validate:"required" example:"Abc@123"`
}

type EmailVerifyReq struct {
	Email string `json:"email" validate:"required" example:"valid_email@folan.com"`
	OTP   string `json:"otp"`
}

func UserRegisterToUserDomain(up *UserRegisterReq) *user.User {
	return &user.User{
		Role:       role.Role{},
		FullName:   &up.FullName,
		Email:      up.Email,
		NationalID: up.NationalID,
		Password:   up.Password,
	}
}
