package handler

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/models"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"github.com/yanarowana123/onelab2/internal/services"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

var router *mux.Router

func TestMain(m *testing.M) {
	dir, err := filepath.Abs(filepath.Join(filepath.Dir("./"), "../../../"))

	err = gotenv.Load(filepath.Join(dir, ".env.test"))
	if err != nil {
		log.Fatal(err)
	}
	config, err := configs.New()

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()

	repositoryManager := repositories.NewManager(*config)

	serviceManager := services.NewManager(*repositoryManager, *config)

	handlerManager := NewManager(*serviceManager, validate)

	router = mux.NewRouter()
	router.HandleFunc("/signup", handlerManager.Register()).Methods("POST")
	//r.HandleFunc("/login", h.LogMiddleware(h.Login())).Methods("POST")
	//r.HandleFunc("/refresh", h.RefreshToken()).Methods("POST")
	//
	//r.HandleFunc("/user/{userID}", h.TokenValidateMiddleware(h.GetUserByID())).Methods("GET")
	//r.HandleFunc("/users/books", h.TokenValidateMiddleware(h.PaginateMiddleware(h.GetUserListWithBooks()))).Methods("GET")
	//r.HandleFunc("/users/book-quantity", h.TokenValidateMiddleware(h.PaginateMiddleware(h.GetUserListWithBooksQuantity()))).Methods("GET")
	//
	//r.HandleFunc("/book", h.TokenValidateMiddleware(h.CreateBook())).Methods("POST")
	//r.HandleFunc("/book/{bookID}", h.TokenValidateMiddleware(h.GetBookByID())).Methods("GET")
	//
	//r.HandleFunc("/checkout/{bookID}", h.TokenValidateMiddleware(h.CheckOut())).Methods("POST")
	//r.HandleFunc("/return/{bookID}", h.TokenValidateMiddleware(h.Return())).Methods("POST")
	m.Run()
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+"q")
	}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	return writer
}

func TestRegister(t *testing.T) {

	tests := []struct {
		name        string
		requestBody models.CreateUserRequest
		want        int
	}{
		{
			name: "user is registered successfully",
			requestBody: models.CreateUserRequest{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email@r.kz",
				Password:  "Password",
			},
			want: http.StatusCreated,
		},
		{
			name: "user registration not valid email provided error",
			requestBody: models.CreateUserRequest{
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email",
				Password:  "Password",
			},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := makeRequest("POST", "/signup", tt.requestBody, false)

			if writer.Code != tt.want {
				t.Errorf("Status got = %v, want %v", writer.Code, tt.want)
			}
		})
	}
}
