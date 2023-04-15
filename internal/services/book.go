package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type IBookService interface {
	Create(ctx context.Context, book models.CreateBookRequest) (*models.BookResponse, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.BookResponse, error)
}

type BookService struct {
	repository repositories.Manager
	utils      IUtilsService
}

func NewBookService(repository repositories.Manager, utils IUtilsService) *BookService {
	return &BookService{
		repository: repository,
		utils:      utils,
	}
}

func (s *BookService) Create(ctx context.Context, book models.CreateBookRequest) (*models.BookResponse, error) {
	book.ID = s.utils.GenerateID()
	return s.repository.Book.Create(ctx, book)
}

func (s *BookService) GetByID(ctx context.Context, ID uuid.UUID) (*models.BookResponse, error) {
	return s.repository.Book.GetByID(ctx, ID)
}
