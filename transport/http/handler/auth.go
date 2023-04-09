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
		json.NewDecoder(r.Body).Decode(&userAttempting)

		user, err := h.service.User.GetByLogin(r.Context(), userAttempting.Login)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAttempting.Password))

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenPair, err := h.service.Auth.GenerateTokenPair(*user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			tokenPair, err := h.service.Auth.GenerateTokenPair(models.AuthUser{ID: userID, Login: user.Login})

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
