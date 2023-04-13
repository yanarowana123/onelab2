package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserRequest struct {
	ID        uuid.UUID
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required" json:"password"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *CreateUserRequest) ToUserResponse() *UserResponse {
	return &UserResponse{ID: m.ID, FirstName: m.FirstName, LastName: m.LastName, Email: m.Email, CreatedAt: time.Now()}
}

type AuthUser struct {
	ID       uuid.UUID
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserWithBook struct {
	User  UserResponse   `json:"user"`
	Books []BookResponse `json:"books"`
}
type UserWithBookQuantity struct {
	User  UserResponse `json:"user"`
	Books int          `json:"books"`
}

type UserWithBookList struct {
	Users []UserWithBook `json:"users"`
}

type UserWithBookQuantityList struct {
	Users []UserWithBookQuantity `json:"users"`
}
