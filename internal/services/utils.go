package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IUtilsService interface {
	GenerateID() uuid.UUID
	HashPassword(password string) ([]byte, error)
}

type UtilsService struct {
}

func (s *UtilsService) GenerateID() uuid.UUID {
	return uuid.New()
}

func (s *UtilsService) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)

}
