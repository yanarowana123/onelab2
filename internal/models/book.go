package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type CreateBookRequest struct {
	ID     uuid.UUID
	Name   string
	Author string
}

type BookResponse struct {
	ID        uuid.UUID
	Name      string
	Author    string
	CreatedAt time.Time
}

//TODO should I store it here??
type NullableBook struct {
	ID        sql.NullString
	Name      sql.NullString
	Author    sql.NullString
	CreatedAt sql.NullTime
}
