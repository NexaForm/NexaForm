package user

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorWhileGeneratingHashPassword   = "faield to hash password"
	ErrorInvalidEmail                  = "invalid email format"
	ErrorInvalidNationalCodeDigits     = "national code must be exactly 10 digits"
	ErrorInvalidNationalCodeOnlyDigits = "national code must contain only digits"
	ErrorInvalidNationalCode           = "invalid national code"
	emailRegex                         = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
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
