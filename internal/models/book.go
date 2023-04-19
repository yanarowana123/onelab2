package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type CreateBookRequest struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `validate:"required" json:"name"`
	AuthorID uuid.UUID `validate:"required" json:"author_id"`
}

type BookResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type BookResponseWithMoneySum struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	Sum       float64   `json:"sum"`
}
type NullableBook struct {
	ID        sql.NullString
	Name      sql.NullString
	AuthorID  sql.NullString
	CreatedAt sql.NullTime
}

func (m *CreateBookRequest) ToBookResponse() *BookResponse {
	return &BookResponse{ID: m.ID, Name: m.Name, AuthorID: m.AuthorID, CreatedAt: time.Now()}
}
