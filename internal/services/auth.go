package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"time"
)

type IAuthService interface {
	GenerateTokenPair(user models.AuthUser) (*models.Tokens, error)
	GetUserID(token *jwt.Token) uuid.UUID
	GetAccessTokenSecret() string
	GetRefreshTokenSecret() string
}

type AuthService struct {
	accessTokenSecret    string
	refreshTokenSecret   string
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

func NewAuthService(accessTokenSecret, refreshTokenSecret string, accessTokenLifetime, refreshTokenLifetime time.Duration) *AuthService {
	return &AuthService{
		accessTokenSecret:    accessTokenSecret,
		refreshTokenSecret:   refreshTokenSecret,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

func (s *AuthService) GenerateTokenPair(user models.AuthUser) (*models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(s.accessTokenLifetime),
		"id":        user.ID,
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.accessTokenSecret))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(s.refreshTokenLifetime),
		"id":        user.ID,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.refreshTokenSecret))

	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *AuthService) GetUserID(token *jwt.Token) uuid.UUID {
	claims, _ := token.Claims.(jwt.MapClaims)
	userID, _ := uuid.Parse(claims["id"].(string))
	return userID
}

func (s *AuthService) GetAccessTokenSecret() string {
	return s.accessTokenSecret
}

func (s *AuthService) GetRefreshTokenSecret() string {
	return s.refreshTokenSecret
}
