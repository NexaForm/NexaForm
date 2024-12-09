package mappers

import (
	"NexaForm/internal/otp"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/fp"
)

func otpEntityToDomain(entity entities.OTP) otp.OTP {
	return otp.OTP{
		ID:        entity.ID,
		UserID:    entity.User.ID,
		OTPCode:   entity.OTPCode,
		OTPExpiry: entity.OTPExpiry,
	}
}

func BatchOtpEntityToDomain(entities []entities.OTP) []otp.OTP {
	return fp.Map(entities, otpEntityToDomain)
}

func OtpDomainToEntity(domain *otp.OTP) *entities.OTP {
	return &entities.OTP{
		UserID:    domain.UserID,
		OTPCode:   domain.OTPCode,
		OTPExpiry: domain.OTPExpiry,
	}
}

func OtpEntityToDomain(entity *entities.OTP) *otp.OTP {
	return &otp.OTP{
		ID:        entity.ID,
		UserID:    entity.User.ID,
		OTPCode:   entity.OTPCode,
		OTPExpiry: entity.OTPExpiry,
	}
}
