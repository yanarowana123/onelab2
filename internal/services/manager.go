package services

import (
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/repositories"
)

type Manager struct {
	User     IUserService
	Book     IBookService
	Auth     IAuthService
	CheckOut ICheckOutService
}

func NewManager(repository repositories.Manager, config configs.Config) *Manager {
	utilsService := &UtilsService{}
	userService := NewUserService(repository, utilsService)
	bookService := NewBookService(repository, utilsService)
	authService := NewAuthService(config.JWTAccessTokenSecret, config.JWTRefreshTokenSecret, config.JWTAccessTokenTTL, config.JWTRefreshTokenTTL)
	checkOutService := NewCheckOutService(repository)
	return &Manager{User: userService, Book: bookService, Auth: authService, CheckOut: checkOutService}
}
