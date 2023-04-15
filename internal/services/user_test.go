package services

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"github.com/yanarowana123/onelab2/mocks/repositories"
	mockServices "github.com/yanarowana123/onelab2/mocks/services"
	"reflect"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.CreateUserRequest
	}
	tests := []struct {
		name                string
		userID              uuid.UUID
		args                args
		prepareGenerateID   func(utilsService *mockServices.MockIUtilsService, ID uuid.UUID)
		prepareHashPassword func(utilsService *mockServices.MockIUtilsService)
		prepareRepository   func(userRepo *mock_repositories.MockIUserRepository, arg args, userResponse *models.UserResponse, err error)
		want                *models.UserResponse
		err                 error
	}{
		{
			name:   "user is created successfully",
			userID: uuid.New(),
			args: args{ctx: context.Background(), user: models.CreateUserRequest{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email",
				Password:  "123",
			}},
			prepareGenerateID: func(utilsService *mockServices.MockIUtilsService, ID uuid.UUID) {
				utilsService.EXPECT().GenerateID().Return(ID)
			},
			prepareHashPassword: func(utilsService *mockServices.MockIUtilsService) {
				utilsService.EXPECT().HashPassword("123").Return([]byte("123"), nil)
			},
			prepareRepository: func(userRepo *mock_repositories.MockIUserRepository, arg args, userResponse *models.UserResponse, err error) {
				userRepo.EXPECT().Create(arg.ctx, arg.user).Return(userResponse, err)
			},
			want: &models.UserResponse{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email",
			},
			err: nil,
		},
		{
			name:   "error is thrown after password hashing",
			userID: uuid.New(),
			args: args{ctx: context.Background(), user: models.CreateUserRequest{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email",
				Password:  "123",
			}},
			prepareGenerateID: func(utilsService *mockServices.MockIUtilsService, ID uuid.UUID) {
				utilsService.EXPECT().GenerateID().Return(ID)
			},
			prepareHashPassword: func(utilsService *mockServices.MockIUtilsService) {
				utilsService.EXPECT().HashPassword("123").Return([]byte("123"), errors.New("password hashing error"))
			},
			prepareRepository: nil,
			want:              nil,
			err:               errors.New("password hashing error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			utilsService := mockServices.NewMockIUtilsService(ctrl)
			userRepo := mock_repositories.NewMockIUserRepository(ctrl)

			tt.args.user.ID = tt.userID

			if tt.prepareGenerateID != nil {
				tt.prepareGenerateID(utilsService, tt.userID)
			}

			if tt.prepareHashPassword != nil {
				tt.prepareHashPassword(utilsService)
			}

			if tt.prepareRepository != nil {
				tt.prepareRepository(userRepo, tt.args, tt.want, tt.err)
			}

			s := NewUserService(repositories.Manager{User: userRepo}, utilsService)

			got, err := s.Create(tt.args.ctx, tt.args.user)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Create() expected error is = %v, want %v", err, tt.err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
