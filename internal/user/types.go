package user

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorWhileGeneratingHashPassword = "faield to hash password"
	ErrorInvalidEmail                = "invalid email format"
	ErrorInvalidNationalCode         = "invalid national code . must be 10 digits"
	emailRegex                       = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	nationalCodeRegex                = regexp.MustCompile(`^\d{10}$`)
)

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

func ValidateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if !emailRegex.MatchString(email) {
		return "", errors.New(ErrorInvalidEmail)
	}
	return strings.ToLower(email), nil
}

func ValidateNationalCode(nationalCode string) (string, error) {
	nationalCode = strings.TrimSpace(nationalCode)
	if !nationalCodeRegex.MatchString(nationalCode) {
		return "", errors.New(ErrorInvalidNationalCode)
	}
	return nationalCode, nil
}
