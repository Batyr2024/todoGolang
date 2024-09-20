package main

import (
	"github.com/Batyr2024/todoGolang/config"
	"github.com/Batyr2024/todoGolang/docs"
	initialize "github.com/Batyr2024/todoGolang/internal/api/init"
	tg "github.com/Batyr2024/todoGolang/internal/api/telegram"
	"github.com/Batyr2024/todoGolang/internal/usecase/telegram"
	"log"
)

// @title ToDo API
// @version 1.0
// @description Description of the ToDO REST API
// @contact.name   Batyr Abdusalamov
// @contact.url    https://twitter.com/
// @contact.email  b.abdusalamov@nv.dunice.net
// @host      localhost:3000
// @BasePath  /tasks
func main() {
	const pathCfg = "/home/dunice/Документы/todoGolang/.env"
	config, err := config.Load(pathCfg)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	docs.Init()
	tg := tg.Init("7536965859:AAFYNHHynpHSjfIf06IydVrS5FMgYxIvZu4")
	server := initialize.Api(config)
	tgBot := telegram.New(tg)
	go tgBot.SendMessage()
	server.Start(string(config.Port))

}
