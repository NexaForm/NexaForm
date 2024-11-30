package user

import (
	"NexaForm/internal/user"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected string
		err      bool
	}{
		{"user@Example.com", "user@example.com", false},
		{"REIHANE@GMAIL.COM", "reihane@gmail.com", false},
		{"reihane@gmail@j.com", "", true},
		{"invalid-email.com", "", true},
	}

	for _, test := range tests {
		result, err := user.ValidateEmail(test.email)
		if (err != nil) != test.err {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}
		if result != test.expected {
			t.Errorf("For email %s, expected: %s, got: %s", test.email, test.expected, result)
		}
	}
}

func TestValidateNationalCode(t *testing.T) {
	tests := []struct {
		code     string
		expected string
		err      bool
	}{
		{"1234567890", "1234567890", false},
		{"12345", "", true},
		{"98765432101", "", true},
		{"2260308821", "2260308821", false},
	}

	for _, test := range tests {
		result, err := user.ValidateNationalCode(test.code)
		if (err != nil) != test.err {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}
		if result != test.expected {
			t.Errorf("For national code %s, expected: %s, got: %s", test.code, test.expected, result)
		}
	}
}

func TestHashAndCheckPassword(t *testing.T) {
	password := "Securepassword@123"
	hashedPassword, err := user.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPass failed: %v", err)
	}

	if !user.CheckPasswordHash(password, hashedPassword) {
		t.Errorf("CheckPasswordHash failed: the password does not match the hash")
	}
}
