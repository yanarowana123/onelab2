package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
)

// CheckOut
// @Summary checkout book
// @Description checkout book
// @Tags checkout
// @Param checkout body models.CreateCheckoutRequest true "body"
// @Security ApiKeyAuth
// @Success 204
// @Failure 404 "book not found"
// @Failure 400 "you already have checked out this book"
// @Router /checkout [post]
func (h *Manager) CheckOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createCheckOutRequest models.CreateCheckoutRequest
		err := json.NewDecoder(r.Body).Decode(&createCheckOutRequest)

		createCheckOutRequest.UserID = r.Context().Value("userID").(uuid.UUID)
		_, err = h.service.Book.GetByID(r.Context(), createCheckOutRequest.BookID)
		if err != nil {
			h.respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = h.validate.Struct(createCheckOutRequest)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		err = h.service.CheckOut.CheckOut(r.Context(), createCheckOutRequest)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Return
// @Summary return book
// @Description return book
// @Tags checkout
// @Param bookID path string true "Book ID (UUID format)"
// @Security ApiKeyAuth
// @Success 204
// @Failure 404 "book not found"
// @Router /return/{bookID} [post]
func (h *Manager) Return() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var returnBookRequset models.ReturnBookRequest
		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])
		_, err = h.service.Book.GetByID(r.Context(), bookID)
		if err != nil {
			h.respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		returnBookRequset.BookID = bookID
		returnBookRequset.UserID = r.Context().Value("userID").(uuid.UUID)

		err = h.validate.Struct(returnBookRequset)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		err = h.service.CheckOut.Return(r.Context(), returnBookRequset)
		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
