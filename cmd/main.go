package main

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"github.com/yanarowana123/onelab2/internal/services"
	"github.com/yanarowana123/onelab2/transport/http"
	"github.com/yanarowana123/onelab2/transport/http/handler"
)

func init() {
	gotenv.Load()
}

func main() {
	config, err := configs.New()
	if err != nil {
		panic(err)
	}

	repositoryManager := repositories.NewManager(*config)

	serviceManager := services.NewManager(*repositoryManager, *config)

	handlerManager := handler.NewManager(*serviceManager)

	r := mux.NewRouter()
	router := http.InitRouter(r, *handlerManager)

	http.InitServer(*config, router)
}
