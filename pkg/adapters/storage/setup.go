package storage

import (
	"NexaForm/config"
	"NexaForm/pkg/adapters/storage/entities"
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresGormConnection initializes a PostgreSQL GORM connection
func NewPostgresGormConnection(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// CreateGenderEnum checks and creates the enum type for gender if it doesn't already exist
func CreateGenderEnum(db *gorm.DB) error {
	sql := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
			CREATE TYPE gender_enum AS ENUM ('Male', 'Female');
		END IF;
	END $$;
	`
	return db.Exec(sql).Error
}

// CreateVisibilityEnum checks and creates the enum type for visibility if it doesn't already exist
func CreateVisibilityEnum(db *gorm.DB) error {
	sql := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'visibility_enum') THEN
			CREATE TYPE visibility_enum AS ENUM ('All', 'Owner_Admin', 'No_One');
		END IF;
	END $$;
	`
	return db.Exec(sql).Error
}

// QuestionTypeEnum checks and creates the enum type for question type if it doesn't already exist
func CreateQuestionTypeEnum(db *gorm.DB) error {
	sql := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'question_type_enum') THEN
			CREATE TYPE question_type_enum AS ENUM ('Poll', 'Text', 'Quiz');
		END IF;
	END $$;
	`
	return db.Exec(sql).Error
}
func Seed(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).FirstOrCreate(&entities.Role{
		ID:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name: "User",
	}, "name = ?", "User").Error
}

// Migrate handles database schema migrations
func Migrate(db *gorm.DB) error {
	// Create the gender_enum type
	if err := CreateGenderEnum(db); err != nil {
		return fmt.Errorf("failed to create gender_enum type: %w", err)
	}
	// Create the visibility_enum type
	if err := CreateVisibilityEnum(db); err != nil {
		return fmt.Errorf("failed to create visibility_enum type: %w", err)
	}
	// Create the question_type_enum type
	if err := CreateQuestionTypeEnum(db); err != nil {
		return fmt.Errorf("failed to create visibility_enum type: %w", err)
	}
	return db.AutoMigrate(
		&entities.User{},
		&entities.OTP{},
		&entities.Survey{},
		&entities.City{},
		&entities.Answer{},
		&entities.Question{},
		&entities.Option{},
		&entities.RenderCondition{},
		&entities.Role{},
		&entities.Permission{},
		&entities.SurveyRole{},
		&entities.SurveyPermission{},
		&entities.SurveyParticipant{},
		&entities.Notification{},
		&entities.Wallet{},
		&entities.WalletTransaction{},
		&entities.AllowedCity{},
		&entities.Attachment{},
	)
}
func AddExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}
