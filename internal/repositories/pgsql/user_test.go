package pgsql

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		ID      uuid.UUID
		want    *models.UserResponse
		wantErr bool
	}{
		{
			name: "user is created successfully",
			args: args{
				ctx: context.Background(),
				user: models.CreateUserRequest{
					FirstName: "FirstName",
					LastName:  "LastName",
					Email:     "Email@r.ru",
					Password:  "Password",
				},
			},
			ID: uuid.New(),
			want: &models.UserResponse{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email@r.ru",
			},
			wantErr: false,
		},
		{
			name: "user is not created because first_name is too long",
			args: args{
				ctx: context.Background(),
				user: models.CreateUserRequest{
					FirstName: "FirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstNameFirstName",
					LastName:  "LastName",
					Email:     "Email@r.ru",
					Password:  "Password",
				},
			},
			ID:      uuid.New(),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db.Exec("delete from users")
		t.Run(tt.name, func(t *testing.T) {
			r := NewUserRepository(db)
			tt.args.user.ID = tt.ID

			_, err := r.Create(tt.args.ctx, tt.args.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := r.GetByID(tt.args.ctx, tt.ID)

			if (err != nil) != tt.wantErr {
				t.Error("User is not created")
				return
			}

			if !tt.wantErr {
				if got.FirstName != tt.want.FirstName {
					t.Errorf("Create() got = %v, want %v", got.FirstName, tt.want.FirstName)
					return
				}

				if got.LastName != tt.want.LastName {
					t.Errorf("Create() got = %v, want %v", got.LastName, tt.want.LastName)
					return
				}
			}
		})
	}
}
