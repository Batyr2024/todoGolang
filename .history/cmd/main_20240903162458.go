package main

import (
	"log"

	"github.com/Batyr2024/todoGolang/pkg/common/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	viper.SetConfigFile("./pkg/common/envs//env")
}