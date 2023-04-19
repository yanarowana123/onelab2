package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateCheckoutRequest struct {
	ID          uuid.UUID `json:"-"`
	UserID      uuid.UUID `validate:"required" json:"-"`
	BookID      uuid.UUID `validate:"required" json:"book_id"`
	MoneyAmount float64   `validate:"required" json:"money_amount"`
}

type ReturnBookRequest struct {
	UserID uuid.UUID `validate:"required" json:"-"`
	BookID uuid.UUID `validate:"required" json:"book_id"`
}

type CheckoutResponse struct {
	UserID       uuid.UUID `json:"user_id"`
	BookID       uuid.UUID `json:"book_id"`
	CheckedOutAt time.Time `json:"checked_out_at"`
	ReturnedAt   time.Time `json:"return_at"`
}
