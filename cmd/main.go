package main

import (
	"github.com/Batyr2024/todoGolang/db"
	"github.com/Batyr2024/todoGolang/internal/api"
	"github.com/Batyr2024/todoGolang/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()
	
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	
	r := gin.Default()
	h := db.Init(dbUrl)

	r.Use(middleware.CorsMiddleware())

	api.RegisterRoutes(r,h)

	r.Run(port)
}