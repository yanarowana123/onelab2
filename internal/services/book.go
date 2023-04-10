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
}

func NewBookService(repository repositories.Manager) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) Create(ctx context.Context, book models.CreateBookRequest) (*models.BookResponse, error) {
	book.ID = uuid.New()
	return s.repository.Book.Create(ctx, book)
}

func (s *BookService) GetByID(ctx context.Context, ID uuid.UUID) (*models.BookResponse, error) {
	return s.repository.Book.GetByID(ctx, ID)
}
