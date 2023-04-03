package repositories

import "github.com/yanarowana123/onelab2/internal/models"

type IUserRepository interface {
	CreateUser(user models.CreateUserReq) (*models.UserResponse, error)
	GetUser(login string) (*models.UserResponse, error)
}

type Manager struct {
	User IUserRepository
}

func NewManager(userRepository IUserRepository) *Manager {
	return &Manager{
		User: userRepository,
	}
}
