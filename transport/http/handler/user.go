package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Manager) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserReq models.CreateUserReq
		err := json.NewDecoder(r.Body).Decode(&createUserReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte(createUserReq.Password), 14)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createUserReq.Password = string(bytes)
		userResponse, err := h.service.User.CreateUser(createUserReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(userResponse)
	}
}

func (h *Manager) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userLogin := params["userLogin"]

		userResponse, err := h.service.User.GetUser(userLogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(userResponse)
	}
}
