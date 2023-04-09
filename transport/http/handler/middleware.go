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
		//reqBody, _ := ioutil.ReadAll(r.Body)
		//
		//h.logger.InfoLogger.Printf("%s %q", r.Method, html.EscapeString(r.URL.Path))
		//h.logger.InfoLogger.Printf("Request body:%s", reqBody)
		next.ServeHTTP(w, r)
	}
}

func (h *Manager) TokenValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
				}

				return []byte(h.service.Auth.GetAccessTokenSecret()), nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if token.Valid {
				ctx := context.WithValue(r.Context(), "userID", h.service.Auth.GetUserID(token))
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		}
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
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
