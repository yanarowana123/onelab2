package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateCheckOutRequest struct {
	UserID uuid.UUID `validate:"required" json:"user_id"`
	BookID uuid.UUID `validate:"required" json:"book_id"`
}

type CheckOutResponse struct {
	UserID       uuid.UUID `json:"user_id"`
	BookID       uuid.UUID `json:"book_id"`
	CheckedOutAt time.Time `json:"checked_out_at"`
	ReturnedAt   time.Time `json:"return_at"`
}
