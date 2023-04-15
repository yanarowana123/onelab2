package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
	mock_repositories "github.com/yanarowana123/onelab2/mocks/repositories"
	"testing"
)

func TestCheckOutService_CheckOut(t *testing.T) {
	type args struct {
		ctx      context.Context
		checkOut models.CreateCheckOutRequest
	}
	tests := []struct {
		name         string
		args         args
		prepareRepos func(repo *mock_repositories.MockICheckOutRepository)
		wantErr      bool
	}{
		{
			name: "book checked out successfully",
			args: args{
				ctx:      context.Background(),
				checkOut: models.CreateCheckOutRequest{},
			},
			prepareRepos: func(repo *mock_repositories.MockICheckOutRepository) {
				repo.EXPECT().HasUserReturnedBook(gomock.Any(), gomock.Any()).Return(true)
				repo.EXPECT().CheckOut(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "cant check out the book because it is already checked out",
			args: args{
				ctx:      context.Background(),
				checkOut: models.CreateCheckOutRequest{},
			},
			prepareRepos: func(repo *mock_repositories.MockICheckOutRepository) {
				repo.EXPECT().HasUserReturnedBook(gomock.Any(), gomock.Any()).Return(false)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			checkOutRepo := mock_repositories.NewMockICheckOutRepository(ctrl)
			repoManager := repositories.Manager{CheckOut: checkOutRepo}
			checkOutService := NewCheckOutService(repoManager)

			tt.prepareRepos(checkOutRepo)

			err := checkOutService.CheckOut(tt.args.ctx, tt.args.checkOut)

			if err != nil && !tt.wantErr {
				t.Errorf("CheckOut() throw error")
			}
		})
	}
}
