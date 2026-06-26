package usecase

import (
	"context"
	"errors"
	"os"
	"service-wedding/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (string, string, error)
	GetMe(ctx context.Context, userID int64) (*domain.User, error)
	Register(ctx context.Context, name, email, password, role string) error
}

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(ur domain.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: ur}
}

type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Generate Access Token
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key"
	}

	accessTokenClaims := &JWTClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // 2 Hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessStr, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshTokenClaims := &JWTClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 7 Days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshStr, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessStr, refreshStr, nil
}

func (u *authUsecase) GetMe(ctx context.Context, userID int64) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, userID)
}

func (u *authUsecase) Register(ctx context.Context, name, email, password, role string) error {
	existing, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("user with this email already exists")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedBytes),
		Role:         role,
	}

	return u.userRepo.Create(ctx, newUser)
}
