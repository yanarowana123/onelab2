package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/yanarowana123/onelab2/internal/services"
)

type Manager struct {
	service  services.Manager
	validate *validator.Validate
}

func NewManager(service services.Manager, validate *validator.Validate) *Manager {
	return &Manager{service: service, validate: validate}
}
