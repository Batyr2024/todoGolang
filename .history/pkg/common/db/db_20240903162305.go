package db

import (
	"log"
	"github.com/Batyr2024/todoGolang/"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB{
	db,err := gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}