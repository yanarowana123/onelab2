package services

import (
	"context"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type ICheckOutService interface {
	CheckOut(ctx context.Context, checkOut models.CreateCheckOutRequest) error
	Return(ctx context.Context, checkOut models.CreateCheckOutRequest) error
}

type CheckOutService struct {
	repository repositories.Manager
}

func NewCheckOutService(repository repositories.Manager) *CheckOutService {
	return &CheckOutService{
		repository: repository,
	}
}

func (s CheckOutService) CheckOut(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	return s.repository.CheckOut.CheckOut(ctx, checkOut)
}

func (s CheckOutService) Return(ctx context.Context, checkOut models.CreateCheckOutRequest) error {
	return s.repository.CheckOut.Return(ctx, checkOut)
}
