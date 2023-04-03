package handler

import (
	"github.com/yanarowana123/onelab2/internal/services"
	"github.com/yanarowana123/onelab2/pkg/logger"
)

type Manager struct {
	logger  logger.Logger
	service services.Manager
}

func NewManager(logger logger.Logger, service services.Manager) *Manager {
	return &Manager{logger, service}
}
