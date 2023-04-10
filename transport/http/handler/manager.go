package handler

import (
	"github.com/yanarowana123/onelab2/internal/services"
)

type Manager struct {
	service services.Manager
}

func NewManager(service services.Manager) *Manager {
	return &Manager{service}
}
