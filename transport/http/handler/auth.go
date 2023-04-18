package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yanarowana123/onelab2/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
// @Summary Registration
// @Description Registration
// @Tags auth
// @Param user body models.CreateUserRequest true "body"
// @Success 200 {object} models.UserResponse
// @Failure 400 "validation error"
// @Router /signup [post]
func (h *Manager) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createUserReq models.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&createUserReq)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.validate.Struct(createUserReq)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		userResp, err := h.service.User.Create(ctx, createUserReq)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userResp)
	}
}

// Login
// @Summary log-in to system
// @Description log-in to system
// @Tags auth
// @Param auth body models.AuthUser true "body"
// @Success 200 {object} models.Tokens
// @Router /login [post]
func (h *Manager) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var userAttempting models.AuthUser
		err := json.NewDecoder(r.Body).Decode(&userAttempting)
		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, err.Error())

			return
		}

		err = h.validate.Struct(userAttempting)
		if err != nil {
			h.respondWithErrorList(w, http.StatusBadRequest, err)
			return
		}

		user, err := h.service.User.GetByLogin(r.Context(), userAttempting.Email)

		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, "email or password is incorrect")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAttempting.Password))

		if err != nil {
			h.respondWithError(w, http.StatusBadRequest, "email or password is incorrect")
			return
		}
		tokenPair, err := h.service.Auth.GenerateTokenPair(*user)

		if err != nil {
			h.respondWithError(w, http.StatusInternalServerError, err.Error())
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
			h.respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if refreshToken.Valid {
			userID := h.service.Auth.GetUserID(refreshToken)
			user, err := h.service.User.GetByID(r.Context(), userID)

			if err != nil {
				h.respondWithError(w, http.StatusUnauthorized, err.Error())
				return
			}

			//TODO add  UserResponse to AuthUser mapper function
			tokenPair, err := h.service.Auth.GenerateTokenPair(models.AuthUser{ID: userID, Email: user.Email})

			if err != nil {
				h.respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			json.NewEncoder(w).Encode(tokenPair)
			return
		} else {
			h.respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

	}
}
