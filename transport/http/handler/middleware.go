package handler

import (
	"html"
	"io/ioutil"
	"net/http"
)

func (h *Manager) LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)

		h.logger.InfoLogger.Printf("%s %q", r.Method, html.EscapeString(r.URL.Path))
		h.logger.InfoLogger.Printf("Request body:%s", reqBody)
		next.ServeHTTP(w, r)

		//TODO log response body
	}
}
