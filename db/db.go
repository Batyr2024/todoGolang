package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	config "github.com/Batyr2024/todoGolang/config"
)

func Connect(cfg config.Config) *gorm.DB {
	db, errCfg := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if errCfg != nil {
		log.Fatalln(errCfg)
	}
	return db
}
