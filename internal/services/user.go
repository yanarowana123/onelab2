package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
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
	utils      IUtilsService
}

func NewUserService(repository repositories.Manager, utils IUtilsService) *UserService {
	return &UserService{
		repository: repository,
		utils:      utils,
	}
}

func (s *UserService) GetListWithBooksQuantity(ctx context.Context, page, pageSize int, dateFrom time.Time) (*models.UserWithBookQuantityList, error) {
	return s.repository.User.GetListWithBooksQuantity(ctx, page, pageSize, dateFrom)
}

func (s *UserService) GetListWithBooks(ctx context.Context, page, pageSize int) (*models.UserWithBookList, error) {
	userWithBookList, err := s.repository.User.GetListWithBooks(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	users := userWithBookList.Users

	booksMoneySum := make(map[uuid.UUID]float64)
	for _, userWithBook := range users {
		books := userWithBook.Books

		for i := range books {
			if sum, ok := booksMoneySum[books[i].ID]; ok {
				books[i].Sum = sum
			} else {
				//TODO handle exception
				sum, err = s.repository.Transaction.GetSumByBookID(ctx, books[i].ID)
				(&books[i]).Sum = sum
				booksMoneySum[books[i].ID] = sum
			}
		}
	}
	return userWithBookList, nil
}

func (s *UserService) Create(ctx context.Context, user models.CreateUserRequest) (*models.UserResponse, error) {
	user.ID = s.utils.GenerateID()

	bytes, err := s.utils.HashPassword(user.Password)
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
