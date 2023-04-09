package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateCheckOutRequest struct {
	UserID uuid.UUID
	BookID uuid.UUID
}

type CheckOutResponse struct {
	UserID       uuid.UUID
	BookID       uuid.UUID
	CheckedOutAt time.Time
	ReturnedAt   time.Time
}
