package main

import (
	"log"

	"github.com/Batyr2024/todoGolang/pkg/common/db"
	"github.com/Batyr2024/todoGolang/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	
	r := gin.Default()
	db.Init(dbUrl)

	
}