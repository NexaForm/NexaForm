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
		wantErr  bool
	}{
		{"2260308821", "2260308821", false}, // Valid code
		{"1234567890", "", true},            // Invalid code
		{"9876543210", "", true},            // Invalid code
		{"0000000000", "", true},            // Invalid code
		{"12345678", "", true},              // Invalid: less than 10 digits
		{"abc1234567", "", true},            // Invalid: non-digit characters
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			result, err := user.ValidateNationalCode(tt.code)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateNationalCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && result != tt.expected {
				t.Errorf("ValidateNationalCode() = %v, want %v", result, tt.expected)
			}
		})
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
