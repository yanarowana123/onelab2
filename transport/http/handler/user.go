package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/yanarowana123/onelab2/internal/models"
	"net/http"
	"time"
)

// GetUserByID
// @Summary get user by id
// @Description get user by id
// @Tags user
// @Param userID path string true "User ID (UUID format)"
// @Security ApiKeyAuth
// @Success 200 {object} models.UserResponse
// @Failure 404 "user not found"
// @Router /user/{userID} [get]
func (h *Manager) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		userID, err := uuid.Parse(params["userID"])

		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		userResponse, err := h.service.User.GetByID(r.Context(), userID)
		if err != nil {
			h.respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		json.NewEncoder(w).Encode(userResponse)
	}
}

// GetUserListWithBooks
// @Summary get user list with books
// @Description get user list with books
// @Tags user
// @Param page query int false "page (pagination)"
// @Param pageSize query int false "page size (pagination)"
// @Security ApiKeyAuth
// @Success 200 {object} models.UserWithBookList
// @Router /users/books [get]
func (h *Manager) GetUserListWithBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page := r.Context().Value("page").(int)
		pageSize := r.Context().Value("pageSize").(int)
		userListWithBooks, err := h.service.User.GetListWithBooks(r.Context(), page, pageSize)
		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(userListWithBooks)
	}
}

// GetUserListWithBooksQuantity
// @Summary get user list with books quantity they have at the moment
// @Description get user list with books quantity
// @Tags user
// @Param page query int false "page (pagination)"
// @Param pageSize query int false "page size (pagination)"
// @Security ApiKeyAuth
// @Success 200 {object} models.UserWithBookQuantityList
// @Router /users/book-quantity [get]
func (h *Manager) GetUserListWithBooksQuantity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page := r.Context().Value("page").(int)
		pageSize := r.Context().Value("pageSize").(int)
		now := time.Now()
		dateFrom := now.AddDate(0, -1, 0)
		userListWithBooks, err := h.service.User.GetListWithBooksQuantity(r.Context(), page, pageSize, dateFrom)

		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(userListWithBooks)
	}
}
