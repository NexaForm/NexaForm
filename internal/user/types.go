package user

import (
	"NexaForm/internal/role"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenderType defines valid gender values
type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

var (
	ErrUserNotFound                    = errors.New("user not found")
	ErrInvalidEmail                    = errors.New("invalid email format")
	ErrInvalidPassword                 = errors.New("invalid password format")
	ErrEmailAlreadyExists              = errors.New("email already exists")
	ErrInvalidAuthentication           = errors.New("email and password doesn't match")
	ErrorWhileGeneratingHashPassword   = "faield to hash password"
	ErrorInvalidEmail                  = "invalid email format"
	ErrorInvalidNationalCodeDigits     = "national code must be exactly 10 digits"
	ErrorInvalidNationalCodeOnlyDigits = "national code must contain only digits"
	ErrorInvalidNationalCode           = "invalid national code"
	emailRegex                         = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type Repo interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	ActivateUser(ctx context.Context, email string) (*User, error)
}

type User struct {
	ID           uuid.UUID
	FullName     *string
	Email        string
	EmailIsValid bool
	Password     string
	NationalID   string
	Role         role.Role
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) PasswordIsValid(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	isMatched := emailRegex.MatchString(email)
	if !isMatched {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePasswordWithFeedback(password string) error {
	tests := []struct {
		pattern string
		message string
	}{
		{".{7,}", "Password must be at least 7 characters long"},
		{"[a-z]", "Password must contain at least one lowercase letter"},
		{"[A-Z]", "Password must contain at least one uppercase letter"},
		{"[0-9]", "Password must contain at least one digit"},
		{"[^\\d\\w]", "Password must contain at least one special character"},
	}

	var errMessages []string

	for _, test := range tests {
		match, _ := regexp.MatchString(test.pattern, password)
		if !match {
			errMessages = append(errMessages, test.message)
		}
	}

	if len(errMessages) > 0 {
		feedback := strings.Join(errMessages, "\n")
		return errors.Join(ErrInvalidPassword, fmt.Errorf(feedback))
	}

	return nil
}

func LowerCaseEmail(email string) string {
	return strings.ToLower(email)
}
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New(ErrorWhileGeneratingHashPassword)
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// func ValidateEmail(email string) (string, error) {
// 	email = strings.TrimSpace(email)
// 	if !emailRegex.MatchString(email) {
// 		return "", errors.New(ErrorInvalidEmail)
// 	}
// 	return strings.ToLower(email), nil
// }

func ValidateNationalCode(code string) (string, error) {
	if len(code) != 10 {
		return "", errors.New(ErrorInvalidNationalCodeDigits)
	}

	if _, err := strconv.Atoi(code); err != nil {
		return "", errors.New(ErrorInvalidNationalCodeOnlyDigits)
	}

	if code == "0000000000" {
		return "", errors.New(ErrorInvalidNationalCode)
	}
	if code == "9876543210" {
		return "", errors.New(ErrorInvalidNationalCode)
	}
	digits := make([]int, 10)
	for i, c := range code {
		digits[i], _ = strconv.Atoi(string(c))
	}

	var B int
	for i := 0; i < 9; i++ {
		B += digits[i] * (10 - i)
	}

	C := B - (B/11)*11
	A := digits[9]

	if C == 0 && A == C {
		return code, nil
	}
	if C == 1 && A == 1 {
		return code, nil
	}
	if C > 1 && A == (11-C) {
		return code, nil
	}

	return "", errors.New(ErrorInvalidNationalCode)
}
