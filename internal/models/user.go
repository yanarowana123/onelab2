package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserRequest struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

type UserResponse struct {
	ID        uuid.UUID
	Name      string
	Login     string
	CreatedAt time.Time
}

type AuthUser struct {
	ID       uuid.UUID
	Login    string
	Password string
}

type UserWithBook struct {
	User  UserResponse
	Books []BookResponse
}
type UserWithBookQuantity struct {
	User  UserResponse
	Books int
}

type UserWithBookList struct {
	Users []UserWithBook
}

type UserWithBookQuantityList struct {
	Users []UserWithBookQuantity
}
