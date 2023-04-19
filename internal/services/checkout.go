package services

import (
	"context"
	"errors"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type ICheckOutService interface {
	CheckOut(ctx context.Context, checkOut models.CreateCheckoutRequest) error
	Return(ctx context.Context, returnBookRequest models.ReturnBookRequest) error
	hasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckoutRequest) bool
}

type CheckOutService struct {
	repository repositories.Manager
	utils      IUtilsService
}

func NewCheckOutService(repository repositories.Manager, utils IUtilsService) *CheckOutService {
	return &CheckOutService{
		repository: repository,
		utils:      utils,
	}
}

func (s *CheckOutService) CheckOut(ctx context.Context, checkOut models.CreateCheckoutRequest) error {
	if s.hasUserReturnedBook(ctx, checkOut) {
		checkOut.ID = s.utils.GenerateID()
		err := s.repository.CheckOut.CheckOut(ctx, checkOut)
		if err == nil {
			createTransactionRequest := models.CreateTransactionRequest{
				UserID:      checkOut.UserID,
				BookID:      checkOut.BookID,
				CheckoutID:  checkOut.ID,
				MoneyAmount: checkOut.MoneyAmount,
			}
			err = s.repository.Transaction.Create(ctx, createTransactionRequest)

			if err != nil {
				//TODO implement rollback mechanism if transaction fails
			}

			return err
		}
		return err
	}

	return errors.New("you already have checked out this book")
}

func (s *CheckOutService) Return(ctx context.Context, returnBookRequest models.ReturnBookRequest) error {
	return s.repository.CheckOut.Return(ctx, returnBookRequest)
}

func (s *CheckOutService) hasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckoutRequest) bool {
	return s.repository.CheckOut.HasUserReturnedBook(ctx, checkOut)
}
