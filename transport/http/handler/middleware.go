package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
)

func (h *Manager) LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

func (h *Manager) TokenValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
				}

				return []byte(h.service.Auth.GetAccessTokenSecret()), nil
			})

			if err != nil {
				h.respondWithError(w, http.StatusUnauthorized, err.Error())
				return
			}

			if token.Valid {
				userId := h.service.Auth.GetUserID(token)
				ctx := context.WithValue(r.Context(), "userID", userId)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				h.respondWithError(w, http.StatusUnauthorized, err.Error())
				return
			}
		}

		h.respondWithError(w, http.StatusUnauthorized, "Invalid Token")
		return
	}
}

func (h *Manager) PaginateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageString := r.URL.Query().Get("page")
		page := 1
		if len(pageString) > 0 {
			page, _ = strconv.Atoi(pageString)
		}

		pageSizeString := r.URL.Query().Get("pageSize")
		pageSize := 20
		if len(pageString) > 0 {
			pageSize, _ = strconv.Atoi(pageSizeString)
		}

		ctx := context.WithValue(r.Context(), "page", page)
		ctx = context.WithValue(ctx, "pageSize", pageSize)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
