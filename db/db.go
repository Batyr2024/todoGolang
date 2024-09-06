package db

import (
	"github.com/Batyr2024/todoGolang/config"
	"github.com/Batyr2024/todoGolang/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Task{})

	return db
}
