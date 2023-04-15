package services

import (
	"context"
	"errors"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type ICheckOutService interface {
	CheckOut(ctx context.Context, checkOut models.CreateCheckOutRequest) error
	Return(ctx context.Context, checkOut models.CreateCheckOutRequest) error
	HasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckOutRequest) bool
}

type CheckOutService struct {
	repository repositories.Manager
}

func NewCheckOutService(repository repositories.Manager) *CheckOutService {
	return &CheckOutService{
		repository: repository,
	}
}

func (s *CheckOutService) CheckOut(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	if s.HasUserReturnedBook(ctx, checkOut) {
		return s.repository.CheckOut.CheckOut(ctx, checkOut)
	}

	return errors.New("you already have checked out this book")
}

func (s *CheckOutService) Return(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	return s.repository.CheckOut.Return(ctx, checkOut)
}

func (s *CheckOutService) HasUserReturnedBook(ctx context.Context, checkOut models.CreateCheckOutRequest) bool {
	return s.repository.CheckOut.HasUserReturnedBook(ctx, checkOut)
}
