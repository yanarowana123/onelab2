package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yanarowana123/onelab2/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Manager) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var userAttempting models.AuthUser
		err := json.NewDecoder(r.Body).Decode(&userAttempting)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		err = h.validate.Struct(userAttempting)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errors := models.NewErrorsCustomFromValidationErrors(err)
			json.NewEncoder(w).Encode(errors)
			return
		}

		user, err := h.service.User.GetByLogin(r.Context(), userAttempting.Email)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAttempting.Password))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}
		tokenPair, err := h.service.Auth.GenerateTokenPair(*user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorCustom{Msg: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(tokenPair)
	}
}

func (h *Manager) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var tokenReqBody struct {
			RefreshToken string `json:"refresh_token"`
		}

		json.NewDecoder(r.Body).Decode(&tokenReqBody)

		refreshToken, err := jwt.Parse(tokenReqBody.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			return []byte(h.service.Auth.GetRefreshTokenSecret()), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if refreshToken.Valid {
			userID := h.service.Auth.GetUserID(refreshToken)
			user, err := h.service.User.GetByID(r.Context(), userID)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			//TODO add  UserResponse to AuthUser mapper function
			tokenPair, err := h.service.Auth.GenerateTokenPair(models.AuthUser{ID: userID, Email: user.Email})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(tokenPair)
			return
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

	}
}
