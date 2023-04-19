package models

import "github.com/google/uuid"

type CreateTransactionRequest struct {
	UserID      uuid.UUID `json:"user_id" validate:"required, uuid"`
	BookID      uuid.UUID `json:"book_id" validate:"required, uuid"`
	CheckoutID  uuid.UUID `json:"checkout_id" validate:"required, uuid"`
	MoneyAmount float64   `json:"money_amount" validate:"required, number"`
}
