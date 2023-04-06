package main

import (
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2/configs"
	"github.com/yanarowana123/onelab2/internal/repositories"
	"github.com/yanarowana123/onelab2/internal/repositories/memory"
	"github.com/yanarowana123/onelab2/internal/services"
	"github.com/yanarowana123/onelab2/pkg/logger"
	"github.com/yanarowana123/onelab2/transport/http"
	"github.com/yanarowana123/onelab2/transport/http/handler"
)

func init() {
	gotenv.Load()
}

func main() {
	loggerManager, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	config := configs.New()
	//sql
	//db, err := mysql.New(*config)
	//defer db.Close()
	//if err != nil {
	//	panic(err)
	//}
	//userRepository := repository.NewMysqlUserRepository(db)

	//in-memory
	userRepository := repository.NewMemoryUserRepository() // это должно быть внутри своего слоя
	repositoryManager := repositories.NewManager(userRepository)

	userService := services.NewUserService(*repositoryManager) // это должно быть внутри своего слоя
	serviceManager := services.NewManager(userService)

	handlerManager := handler.NewManager(*loggerManager, *serviceManager)

	router := http.InitRouter(*handlerManager) // роутер можно вынести в слой транспорта

	http.InitServer(*config, router)
}
