package initialize

import (
	config "github.com/Batyr2024/todoGolang/config"
	db "github.com/Batyr2024/todoGolang/db"
	http "github.com/Batyr2024/todoGolang/internal/api"
	handler "github.com/Batyr2024/todoGolang/internal/api/handler"
	repository "github.com/Batyr2024/todoGolang/internal/repository"
	usecase "github.com/Batyr2024/todoGolang/internal/usecase"
)

func Api(cfg config.Config) *http.ServerHTTP {
	//wire.Build(db.Connect, repository.New, usecase.New, handler.New, http.NewServer)
	gormDB := db.Connect(cfg)

	taskRepository := repository.New(gormDB)
	taskUseCase := usecase.New(taskRepository)
	taskHandler := handler.New(taskUseCase)
	serverHTTP := http.NewServer(taskHandler)
	return serverHTTP
}
