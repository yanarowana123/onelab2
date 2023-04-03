package services

import (
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type IUserService interface {
	CreateUser(user models.CreateUserReq) (*models.UserResponse, error)
	GetUser(login string) (*models.UserResponse, error)
}

type UserService struct {
	repository repositories.Manager
}

func NewUserService(repository repositories.Manager) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(user models.CreateUserReq) (*models.UserResponse, error) {
	return s.repository.User.CreateUser(user)
}

func (s *UserService) GetUser(login string) (*models.UserResponse, error) {
	return s.repository.User.GetUser(login)
}
