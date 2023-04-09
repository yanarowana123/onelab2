package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
	"time"
)

func (h *Manager) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserReq models.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&createUserReq)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		userResp, err := h.service.User.Create(ctx, createUserReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userResp)
	}
}

func (h *Manager) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID, err := uuid.Parse(params["userID"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userResponse, err := h.service.User.GetByID(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(userListWithBooks)
	}
}
