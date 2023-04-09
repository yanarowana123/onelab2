package pgsql

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db,
	}
}

func (r BookRepository) Create(ctx context.Context, book models.CreateBookRequest) (*models.BookResponse, error) {
	_, err := r.db.ExecContext(ctx, "insert into books (id, name, author) values ($1,$2,$3)",
		book.ID, book.Name, book.Author)
	if err != nil {
		return nil, err
	}

	return &models.BookResponse{ID: book.ID, Name: book.Name, Author: book.Author}, nil
}

func (r BookRepository) GetByID(ctx context.Context, ID uuid.UUID) (*models.BookResponse, error) {
	var bookResponse models.BookResponse
	err := r.db.QueryRowContext(ctx, "SELECT id, name,author, created_at FROM books WHERE id=$1", ID).
		Scan(&bookResponse.ID, &bookResponse.Name, &bookResponse.Author, &bookResponse.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &bookResponse, nil
}
