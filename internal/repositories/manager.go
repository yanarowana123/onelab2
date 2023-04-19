package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories/integration"
	"github.com/yanarowana123/onelab2/internal/repositories/pgsql"
	"time"
)

type IUserRepository interface {
	Create(ctx context.Context, user models.CreateUserRequest) (*models.UserResponse, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.UserResponse, error)
	GetByLogin(ctx context.Context, login string) (*models.AuthUser, error)
	GetListWithBooks(ctx context.Context, page, pageSize int) (*models.UserWithBookList, error)
	GetListWithBooksQuantity(ctx context.Context, page, pageSize int, dateFrom time.Time) (*models.UserWithBookQuantityList, error)
}

type IBookRepository interface {
	Create(ctx context.Context, book models.CreateBookRequest) (*models.BookResponse, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.BookResponse, error)
}

type ICheckoutRepository interface {
	CheckOut(ctx context.Context, checkout models.CreateCheckoutRequest) error
	Return(ctx context.Context, returnBookRequest models.ReturnBookRequest) error
	HasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckoutRequest) bool
}

type ITransactionRepository interface {
	Create(ctx context.Context, createTransactionRequest models.CreateTransactionRequest) error
	GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error)
}

type Manager struct {
	User        IUserRepository
	Book        IBookRepository
	CheckOut    ICheckoutRepository
	Transaction ITransactionRepository
}

func NewManager(config configs.Config) *Manager {
	connection := pgsql.ConnectDB(config.PgSqlDSN)
	userRepository := pgsql.NewUserRepository(connection)
	bookRepository := pgsql.NewBookRepository(connection)
	checkOutRepository := pgsql.NewCheckoutRepository(connection)
	transactionRepository := integration.NewTransactionRepository(config.TransactionBaseUrl, config.TransactionHTTPTimeout, config.TransactionSecretKey)

	return &Manager{
		User:        userRepository,
		Book:        bookRepository,
		CheckOut:    checkOutRepository,
		Transaction: transactionRepository,
	}
}
