package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	config "github.com/Batyr2024/todoGolang/config"
	migration "github.com/Batyr2024/todoGolang/db/migration"
)

func Connect(cfg config.Config) *gorm.DB {
	db, errCfg := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if errCfg != nil {
		log.Fatalln(errCfg)
	}

	errMigrate := migration.UpTask(db)
	if errMigrate!=nil{
		log.Fatalln(errMigrate)
	}

	return db
}
