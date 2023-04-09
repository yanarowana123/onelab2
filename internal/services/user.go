package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type IUserService interface {
	Create(ctx context.Context, user models.CreateUserRequest) (*models.UserResponse, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.UserResponse, error)
	GetByLogin(ctx context.Context, login string) (*models.AuthUser, error)
	GetListWithBooks(ctx context.Context, page, pageSize int) (*models.UserWithBookList, error)
	GetListWithBooksQuantity(ctx context.Context, page, pageSize int, dateFrom time.Time) (*models.UserWithBookQuantityList, error)
}

type UserService struct {
	repository repositories.Manager
}

func NewUserService(repository repositories.Manager) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetListWithBooksQuantity(ctx context.Context, page, pageSize int, dateFrom time.Time) (*models.UserWithBookQuantityList, error) {
	return s.repository.User.GetListWithBooksQuantity(ctx, page, pageSize, dateFrom)
}

func (s *UserService) GetListWithBooks(ctx context.Context, page, pageSize int) (*models.UserWithBookList, error) {
	return s.repository.User.GetListWithBooks(ctx, page, pageSize)
}

func (s *UserService) Create(ctx context.Context, user models.CreateUserRequest) (*models.UserResponse, error) {
	user.ID = uuid.New()

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	user.Password = string(bytes)
	return s.repository.User.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, ID uuid.UUID) (*models.UserResponse, error) {
	return s.repository.User.GetByID(ctx, ID)
}

func (s *UserService) GetByLogin(ctx context.Context, login string) (*models.AuthUser, error) {
	return s.repository.User.GetByLogin(ctx, login)
}
