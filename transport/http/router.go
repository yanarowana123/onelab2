package http

import (
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/transport/http/handler"
)

func InitRouter(h handler.Manager) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", h.LogMiddleware(h.CreateUser())).Methods("POST")
	r.HandleFunc("/user/{userLogin}", h.LogMiddleware(h.GetUser())).Methods("GET")

	return r
}
