package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
	"time"
)

func (h *Manager) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createUserReq models.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&createUserReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		err = h.validate.Struct(createUserReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errors := models.NewErrorsCustomFromValidationErrors(err)
			json.NewEncoder(w).Encode(errors)
			return
		}

		ctx := r.Context()
		userResp, err := h.service.User.Create(ctx, createUserReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userResp)
	}
}

func (h *Manager) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		userID, err := uuid.Parse(params["userID"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		userResponse, err := h.service.User.GetByID(r.Context(), userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(userResponse)
	}
}

func (h *Manager) GetUserListWithBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page := r.Context().Value("page").(int)
		pageSize := r.Context().Value("pageSize").(int)
		userListWithBooks, err := h.service.User.GetListWithBooks(r.Context(), page, pageSize)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(userListWithBooks)
	}
}

func (h *Manager) GetUserListWithBooksQuantity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page := r.Context().Value("page").(int)
		pageSize := r.Context().Value("pageSize").(int)
		now := time.Now()
		dateFrom := now.AddDate(0, -1, 0)
		userListWithBooks, err := h.service.User.GetListWithBooksQuantity(r.Context(), page, pageSize, dateFrom)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(userListWithBooks)
	}
}
