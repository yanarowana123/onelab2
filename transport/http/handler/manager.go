package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/services"
	"net/http"
)

type Manager struct {
	service  services.Manager
	validate *validator.Validate
}

func NewManager(service services.Manager, validate *validator.Validate) *Manager {
	return &Manager{service: service, validate: validate}
}

func (h *Manager) respondWithError(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorCustom{Msg: errorMsg})
}

func (h *Manager) respondWithErrorList(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	errors := models.NewErrorsCustomFromValidationErrors(err)
	json.NewEncoder(w).Encode(errors)
}
