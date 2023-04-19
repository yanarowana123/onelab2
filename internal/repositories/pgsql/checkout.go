package pgsql

import (
	"context"
	"database/sql"
	"github.com/yanarowana123/onelab2/internal/models"
)

type CheckoutRepository struct {
	db *sql.DB
}

func NewCheckoutRepository(db *sql.DB) *CheckoutRepository {
	return &CheckoutRepository{
		db,
	}
}

func (r CheckoutRepository) CheckOut(ctx context.Context, checkOut models.CreateCheckoutRequest) error {
	_, err := r.db.ExecContext(ctx, "insert into checkout (id, user_id, book_id) values ($1, $2, $3)",
		checkOut.ID, checkOut.UserID, checkOut.BookID)
	return err
}

func (r CheckoutRepository) Return(ctx context.Context, returnBookRequest models.ReturnBookRequest) error {
	_, err := r.db.ExecContext(ctx, "update checkout set returned_at = now() where user_id = $1 and book_id = $2 and returned_at is null",
		returnBookRequest.UserID, returnBookRequest.BookID)
	return err
}

func (r CheckoutRepository) HasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckoutRequest) bool {
	err := r.db.QueryRowContext(ctx,
		"SELECT user_id from checkout where user_id = $1 and book_id = $2 and returned_at is null order by checked_out_at DESC limit 1", checkOut.UserID, checkOut.BookID).Scan()
	if err == sql.ErrNoRows {
		return true
	}

	return false
}
