package main

import (
	"github.com/Batyr2024/todoGolang/config"
	"github.com/Batyr2024/todoGolang/docs"
	initialize "github.com/Batyr2024/todoGolang/internal/api/init"
	"log"
)

// @title ToDo API
// @version 1.0
// @description Description of the ToDO REST API
// @contact.name   BatyrAbdusalamov
// @contact.url    https://twitter.com/
// @contact.email  b.abdusalamov@nv.dunice.net
// @host      localhost:3000
// @BasePath  /tasks
func main() {
	const pathCfg = "/home/dunice/Документы/todoGolang/.env"
	cfg, err := config.Load(pathCfg)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	docs.Init()

	server := initialize.Api(cfg)

	server.Start(string(cfg.Port))

}
