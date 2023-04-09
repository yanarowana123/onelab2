package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"time"
)

type IAuthService interface {
	GenerateTokenPair(user models.AuthUser) (map[string]string, error)
	GetUserID(token *jwt.Token) uuid.UUID
	GetAccessTokenSecret() string
	GetRefreshTokenSecret() string
}

type AuthService struct {
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthService(accessTokenSecret, refreshTokenSecret string) *AuthService {
	return &AuthService{
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
	}
}

func (s *AuthService) GenerateTokenPair(user models.AuthUser) (map[string]string, error) {

	//TODO read ExpiresAt from config
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(10 * time.Minute),
		"id":        user.ID,
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.accessTokenSecret))

	if err != nil {
		return nil, err
	}

	//TODO read ExpiresAt from config
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(24 * time.Hour),
		"id":        user.ID,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.refreshTokenSecret))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessToken":  accessTokenString,
		"refreshToken": refreshTokenString,
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
