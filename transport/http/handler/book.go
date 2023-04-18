package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
)

// CreateBook
// @Summary create book
// @Description create book
// @Tags book
// @Param book body models.CreateBookRequest true "body"
// @Success 200 {object} models.BookResponse
// @Router /book [post]
func (h *Manager) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createBookRequest models.CreateBookRequest
		err := json.NewDecoder(r.Body).Decode(&createBookRequest)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.validate.Struct(createBookRequest)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		bookResponse, err := h.service.Book.Create(ctx, createBookRequest)
		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookResponse)
	}
}

// GetBookByID
// @Summary get book by id
// @Description get book by id
// @Tags book
// @Param bookID path string true "Book ID (UUID format)"
// @Security ApiKeyAuth
// @Success 200 {object} models.BookResponse
// @Failure 404 "book not found"
// @Router /book/{bookID} [get]
func (h *Manager) GetBookByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])

		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := r.Context()
		bookResponse, err := h.service.Book.GetByID(ctx, bookID)
		if err != nil {
			h.respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		json.NewEncoder(w).Encode(bookResponse)
	}
}
