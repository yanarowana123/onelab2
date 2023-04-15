package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"github.com/yanarowana123/onelab2/internal/services"
	"github.com/yanarowana123/onelab2/transport/http"
	"github.com/yanarowana123/onelab2/transport/http/handler"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	err := gotenv.Load()
	if err != nil {
		return err
	}

	config, err := configs.New()
	if err != nil {
		return err
	}

	validate := validator.New()

	repositoryManager := repositories.NewManager(*config)

	serviceManager := services.NewManager(*repositoryManager, *config)

	handlerManager := handler.NewManager(*serviceManager, validate)

	r := mux.NewRouter()
	router := http.InitRouter(r, *handlerManager)

	http.InitServer(*config, router)
	return nil
}
