package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
	mock_repositories "github.com/yanarowana123/onelab2/mocks/repositories"
	mock_services "github.com/yanarowana123/onelab2/mocks/services"
	"testing"
)

func TestCheckOutService_CheckOut(t *testing.T) {
	type args struct {
		ctx      context.Context
		checkOut models.CreateCheckoutRequest
	}
	tests := []struct {
		name                   string
		args                   args
		prepareCheckoutRepo    func(repo *mock_repositories.MockICheckoutRepository)
		prepareTransactionRepo func(repo *mock_repositories.MockITransactionRepository)
		prepareUtils           func(service *mock_services.MockIUtilsService)
		wantErr                bool
	}{
		{
			name: "book checked out successfully",
			args: args{
				ctx:      context.Background(),
				checkOut: models.CreateCheckoutRequest{},
			},
			prepareCheckoutRepo: func(repo *mock_repositories.MockICheckoutRepository) {
				repo.EXPECT().HasUserReturnedBook(gomock.Any(), gomock.Any()).Return(true)
				repo.EXPECT().CheckOut(gomock.Any(), gomock.Any()).Return(nil)
			},
			prepareTransactionRepo: func(repo *mock_repositories.MockITransactionRepository) {
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			prepareUtils: func(service *mock_services.MockIUtilsService) {
				service.EXPECT().GenerateID().Return(uuid.New())
			},
			wantErr: false,
		},
		{
			name: "cant check out the book because it is already checked out",
			args: args{
				ctx:      context.Background(),
				checkOut: models.CreateCheckoutRequest{},
			},
			prepareCheckoutRepo: func(repo *mock_repositories.MockICheckoutRepository) {
				repo.EXPECT().HasUserReturnedBook(gomock.Any(), gomock.Any()).Return(false)
			},
			prepareTransactionRepo: nil,
			prepareUtils:           nil,
			wantErr:                true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			checkOutRepo := mock_repositories.NewMockICheckoutRepository(ctrl)
			transactionRepo := mock_repositories.NewMockITransactionRepository(ctrl)
			repoManager := repositories.Manager{CheckOut: checkOutRepo, Transaction: transactionRepo}
			utilsService := mock_services.NewMockIUtilsService(ctrl)
			checkOutService := NewCheckOutService(repoManager, utilsService)

			if tt.prepareCheckoutRepo != nil {
				tt.prepareCheckoutRepo(checkOutRepo)

			}
			if tt.prepareTransactionRepo != nil {
				tt.prepareTransactionRepo(transactionRepo)
			}
			if tt.prepareUtils != nil {
				tt.prepareUtils(utilsService)
			}

			err := checkOutService.CheckOut(tt.args.ctx, tt.args.checkOut)

			if err != nil && !tt.wantErr {
				t.Errorf("CheckOut() throw error")
			}
		})
	}
}
