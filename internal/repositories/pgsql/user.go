package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) GetListWithBooksQuantity(ctx context.Context, page, pageSize int, dateFrom time.Time) (*models.UserWithBookQuantityList, error) {
	offset := pageSize * (page - 1)
	rows, err := r.db.QueryContext(ctx, "SELECT u.id, u.email, u.first_name, u.last_name, u.created_at, COUNT(c.user_id)  FROM users u LEFT JOIN check_out c ON u.id = c.user_id AND c.checked_out_at >= $1 GROUP BY u.id, u.first_name, u.email, u.id, u.last_name, u.created_at, u.id, u.created_at ORDER BY u.created_at OFFSET $2 LIMIT $3",
		dateFrom, offset, pageSize)

	if err != nil {
		return nil, err
	}

	var res models.UserWithBookQuantityList

	for rows.Next() {
		var userWithBookQuantity models.UserWithBookQuantity
		err = rows.Scan(
			&userWithBookQuantity.User.ID,
			&userWithBookQuantity.User.Email,
			&userWithBookQuantity.User.FirstName,
			&userWithBookQuantity.User.LastName,
			&userWithBookQuantity.User.CreatedAt,
			&userWithBookQuantity.Books,
		)
		res.Users = append(res.Users, userWithBookQuantity)
	}

	return &res, nil
}

func (r *UserRepository) GetListWithBooks(ctx context.Context, page, pageSize int) (*models.UserWithBookList, error) {
	offset := pageSize * (page - 1)
	rows, err := r.db.QueryContext(ctx, "SELECT u.id, u.email, u.first_name, u.last_name, u.created_at, b.id, b.name, b.author_id, b.created_at  FROM users u LEFT JOIN check_out c ON u.id = c.user_id AND c.returned_at IS NULL LEFT JOIN books b ON c.book_id = b.id ORDER BY u.created_at OFFSET $1 LIMIT $2",
		offset, pageSize)

	if err != nil {
		return nil, err
	}

	var res models.UserWithBookList

	var userWithBookList []models.UserWithBook

	userBookMap := make(map[models.UserResponse][]models.BookResponse)

	for rows.Next() {
		var nullableBook models.NullableBook
		var user models.UserResponse
		var book models.BookResponse
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&nullableBook.ID,
			&nullableBook.Name,
			&nullableBook.AuthorID,
			&nullableBook.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if nullableBook.ID.Valid {
			book.ID, _ = uuid.Parse(nullableBook.ID.String)
			book.Name = nullableBook.Name.String
			book.AuthorID, _ = uuid.Parse(nullableBook.AuthorID.String)
			book.CreatedAt = nullableBook.CreatedAt.Time
			userBookMap[user] = append(userBookMap[user], book)
		} else {
			userBookMap[user] = []models.BookResponse{}
		}
	}

	for user, bookList := range userBookMap {
		var userWithBook models.UserWithBook
		userWithBook.User = user
		userWithBook.Books = bookList
		userWithBookList = append(userWithBookList, userWithBook)
	}

	res.Users = userWithBookList

	return &res, nil
}

func (r *UserRepository) Create(ctx context.Context, user models.CreateUserRequest) (*models.UserResponse, error) {
	_, err := r.db.ExecContext(ctx, "insert into users (id, email, first_name, last_name, password) values ($1,$2,$3,$4, $5)",
		user.ID, user.Email, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return nil, err
	}

	return user.ToUserResponse(), nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID uuid.UUID) (*models.UserResponse, error) {
	var userResponse models.UserResponse
	err := r.db.QueryRowContext(ctx, "SELECT id, email, first_name, last_name, created_at FROM users WHERE id=$1", ID).
		Scan(&userResponse.ID, &userResponse.Email, &userResponse.FirstName, &userResponse.LastName, &userResponse.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &userResponse, nil
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*models.AuthUser, error) {
	var userResponse models.AuthUser
	err := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email=$1", login).Scan(&userResponse.ID, &userResponse.Email, &userResponse.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &userResponse, nil
}
