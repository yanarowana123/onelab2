package http

import (
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2/transport/http/handler"
)

func InitRouter(r *mux.Router, h handler.Manager) *mux.Router {
	r.HandleFunc("/signup", h.LogMiddleware(h.CreateUser())).Methods("POST")
	r.HandleFunc("/login", h.LogMiddleware(h.Login())).Methods("POST")
	r.HandleFunc("/refresh", h.RefreshToken()).Methods("POST")

	r.HandleFunc("/user/{userID}", h.TokenValidateMiddleware(h.GetUserByID())).Methods("GET")
	r.HandleFunc("/users/books", h.TokenValidateMiddleware(h.PaginateMiddleware(h.GetUserListWithBooks()))).Methods("GET")
	r.HandleFunc("/users/book-quantity", h.TokenValidateMiddleware(h.PaginateMiddleware(h.GetUserListWithBooksQuantity()))).Methods("GET")

	r.HandleFunc("/book", h.TokenValidateMiddleware(h.CreateBook())).Methods("POST")
	r.HandleFunc("/book/{bookID}", h.TokenValidateMiddleware(h.GetBookByID())).Methods("GET")

	r.HandleFunc("/checkout/{bookID}", h.TokenValidateMiddleware(h.CheckOut())).Methods("POST")
	r.HandleFunc("/return/{bookID}", h.TokenValidateMiddleware(h.Return())).Methods("POST")

	return r
}
