package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yanarowana123/onelab2/internal/models"
)

type CheckOutRepository struct {
	db *sql.DB
}

func NewCheckOutRepository(db *sql.DB) *CheckOutRepository {
	return &CheckOutRepository{
		db,
	}
}

func (r CheckOutRepository) CheckOut(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	if r.userHasReturnedBook(ctx, checkOut) == false {
		return errors.New("you already have checked out this book")
	}

	_, err := r.db.ExecContext(ctx, "insert into check_out (user_id, book_id) values ($1,$2)",
		checkOut.UserID, checkOut.BookID)
	if err != nil {
		return err
	}

	return nil
}

func (r CheckOutRepository) Return(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	_, err := r.db.ExecContext(ctx, "update check_out set returned_at = now() where user_id = $1 and book_id = $2 and returned_at is null",
		checkOut.UserID, checkOut.BookID)
	return err
}

func (r CheckOutRepository) userHasReturnedBook(ctx context.Context, checkOut models.CreateCheckOutRequest) bool {
	err := r.db.QueryRowContext(ctx,
		"SELECT user_id from check_out where user_id = $1 and book_id = $2 and returned_at is null order by checked_out_at DESC limit 1", checkOut.UserID, checkOut.BookID).Scan()
	if err == sql.ErrNoRows {
		return true
	}

	return false
}
