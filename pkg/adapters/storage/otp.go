package storage

import (
	"NexaForm/internal/otp"
	"NexaForm/pkg/adapters/storage/entities"
	"NexaForm/pkg/adapters/storage/mappers"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type otpRepo struct {
	db *gorm.DB
}

func NewOtpRepo(db *gorm.DB) otp.Repo {
	return &otpRepo{
		db: db,
	}
}
func (r *otpRepo) Create(ctx context.Context, otp *otp.OTP) (*otp.OTP, error) {
	newOtp := mappers.OtpDomainToEntity(otp)
	if err := r.db.WithContext(ctx).Create(&newOtp).Error; err != nil {
		return nil, err
	}
	createdOtp := mappers.OtpEntityToDomain(newOtp)
	return createdOtp, nil
}
func (r *otpRepo) GetByUserIdAndCode(ctx context.Context, userId uuid.UUID, code string) (*otp.OTP, error) {
	var o entities.OTP
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND otp_code = ?", userId, code).
		First(&o).Error
	if err != nil {
		return nil, err
	}
	retreivedOtp := mappers.OtpEntityToDomain(&o)
	return retreivedOtp, nil
}
