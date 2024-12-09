package service

import (
	"NexaForm/internal/otp"
	"NexaForm/internal/user"
	"NexaForm/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"time"

	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

var (
	ErrCreateOtpFaild         error = errors.New("faild to generate or sent otp code")
	ErrEmailVerificationFaild error = errors.New("email verification faild")
	ErrEmailNotVerified       error = errors.New("email is not verified")
)

type AuthService struct {
	otpOps                 *otp.Ops
	userOps                *user.Ops
	secret                 []byte
	tokenExpiration        uint
	refreshTokenExpiration uint
}

func NewAuthService(otpOps *otp.Ops, userOps *user.Ops, secret []byte,
	tokenExpiration uint, refreshTokenExpiration uint) *AuthService {
	return &AuthService{
		otpOps:                 otpOps,
		userOps:                userOps,
		secret:                 secret,
		tokenExpiration:        tokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}
}

type UserToken struct {
	AuthorizationToken string
	RefreshToken       string
	ExpiresAt          int64
}

func (s *AuthService) CreateUserAndSentOtp(ctx context.Context, user *user.User) (*user.User, error) {
	createdUser, err := s.userOps.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Generate OTP for the newly registered user
	err = s.generateAndSendOTP(ctx, createdUser.ID, createdUser.Email)
	if err != nil {
		return nil, ErrCreateOtpFaild
	}

	return createdUser, nil
}
func (s *AuthService) VerifyEmail(ctx context.Context, email, code string) (*user.User, error) {
	user, err := s.userOps.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	_, err = s.otpOps.GetByUserIdAndCode(ctx, user.ID, code)
	if err != nil {
		return nil, ErrEmailVerificationFaild
	}
	_, err = s.userOps.ActivateUser(ctx, user.Email)
	if err != nil {
		return nil, ErrEmailVerificationFaild
	}
	return user, nil
}
func (s *AuthService) Login(ctx context.Context, email, pass string) (*UserToken, error) {
	fetchedUser, err := s.userOps.GetUserByEmailAndPassword(ctx, email, pass)
	if err != nil {
		return nil, err
	}
	if !fetchedUser.EmailIsValid {
		return nil, ErrEmailNotVerified
	}
	// calc expiration time values
	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(s.tokenExpiration))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.refreshTokenExpiration))
	)

	authToken, err := jwt.CreateToken(s.secret, s.userClaims(fetchedUser, authExp))
	if err != nil {
		return nil, err // todo
	}

	refreshToken, err := jwt.CreateToken(s.secret, s.userClaims(fetchedUser, refreshExp))
	if err != nil {
		return nil, err // todo
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          authExp.Unix(),
	}, nil
}

func (s *AuthService) RefreshAuth(ctx context.Context, refreshToken string) (*UserToken, error) {
	claim, err := jwt.ParseToken(refreshToken, s.secret)
	if err != nil {
		return nil, err
	}

	u, err := s.userOps.GetUserByID(ctx, claim.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, user.ErrUserNotFound
	}

	// calc expiration time values
	var (
		authExp = time.Now().Add(time.Minute * time.Duration(s.tokenExpiration))
	)

	authToken, err := jwt.CreateToken(s.secret, s.userClaims(u, authExp))
	if err != nil {
		return nil, err // todo
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          authExp.UnixMilli(),
	}, nil
}

func (s *AuthService) userClaims(user *user.User, exp time.Time) *jwt.UserClaims {
	return &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: &jwt2.NumericDate{
				Time: exp,
			},
		},
		UserID: user.ID,
		Role:   user.Role.ID,
	}
}

func sendOTP(email, otp string) error {
	// SMTP Configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpEmail := "nexaformquera@gmail.com"
	smtpPassword := "queu vidm vmmn ugwy" // Replace with your app password

	// Email Content
	subject := "Your OTP Code"
	body := fmt.Sprintf("Your OTP code is: %s\nThis code will expire in 5 minutes.", otp)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Send Email
	d := gomail.NewDialer(smtpHost, smtpPort, smtpEmail, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return err
	}

	fmt.Println("OTP sent successfully to:", email)
	return nil
}

// GenerateAndSendOTP generates a 6-digit OTP and stores it in the database
func (s *AuthService) generateAndSendOTP(ctx context.Context, userID uuid.UUID, email string) error {
	otpCode, err := Generate6DigitOTP()
	if err != nil {
		return err
	}
	expiry := time.Now().Add(5 * time.Minute)
	newOtp := otp.OTP{UserID: userID, OTPCode: otpCode, OTPExpiry: expiry}
	_, err = s.otpOps.Create(ctx, &newOtp)
	if err != nil {
		return err
	}

	err = sendOTP(email, otpCode)
	return err
}
