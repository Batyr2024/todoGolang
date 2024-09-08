package migration

import (
	"github.com/Batyr2024/todoGolang/domain"
	"gorm.io/gorm"
)

func UpTask(db *gorm.DB)error{
	if hasTable:= db.Migrator().HasTable(&domain.Task{}); !hasTable{
	err := db.Migrator().CreateTable(&domain.Task{});
	return err
	}
	return nil;
}

func DownTask(db *gorm.DB) error{
	err := db.Migrator().DropTable(&domain.Task{})
	return err
}