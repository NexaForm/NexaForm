package service

import (
	"crypto/rand"
	"fmt"
)

// Generate6DigitOTP generates a random 6-digit OTP
func Generate6DigitOTP() (string, error) {
	otp := make([]byte, 3) // 3 bytes for 6 digits
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}
	// Convert to 6-digit number
	return fmt.Sprintf("%06d", (int(otp[0])<<16|int(otp[1])<<8|int(otp[2]))%1000000), nil
}
