package main

import (
	"github.com/Batyr2024/todoGolang/config"
	initialize "github.com/Batyr2024/todoGolang/internal/api/init"
	"log"
)

func main() {
	const pathCfg = "/home/dunice/Документы/todoGolang/.env"
	config, err := config.Load(pathCfg)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server := initialize.Api(config)

	server.Start(string(config.Port))

}
