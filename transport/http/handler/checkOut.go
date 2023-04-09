package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"net/http"
)

func (h *Manager) CheckOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createCheckOutRequest models.CreateCheckOutRequest
		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])
		createCheckOutRequest.BookID = bookID
		fmt.Println(r.Context().Value("userID"))
		createCheckOutRequest.UserID = r.Context().Value("userID").(uuid.UUID)

		err = h.service.CheckOut.CheckOut(r.Context(), createCheckOutRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *Manager) Return() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createCheckOutRequest models.CreateCheckOutRequest
		params := mux.Vars(r)
		bookID, err := uuid.Parse(params["bookID"])
		createCheckOutRequest.BookID = bookID
		createCheckOutRequest.UserID = r.Context().Value("userID").(uuid.UUID)

		err = h.service.CheckOut.Return(r.Context(), createCheckOutRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
