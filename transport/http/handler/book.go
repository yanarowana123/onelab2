package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
)

func (h *Manager) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createBookRequest models.CreateBookRequest
		err := json.NewDecoder(r.Body).Decode(&createBookRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		err = h.validate.Struct(createBookRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errors := models.NewErrorsCustomFromValidationErrors(err)
			json.NewEncoder(w).Encode(errors)
			return
		}

		ctx := r.Context()
		bookResponse, err := h.service.Book.Create(ctx, createBookRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookResponse)
	}
}

func (h *Manager) GetBookByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		ctx := r.Context()
		bookResponse, err := h.service.Book.GetByID(ctx, bookID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(bookResponse)
	}
}
