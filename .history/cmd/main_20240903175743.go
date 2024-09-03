package main

import (
	"github.com/Batyr2024/todoGolang/pkg/common/db"
	"github.com/Batyr2024/todoGolang/pkg/tasks"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	
	r := gin.Default()
	h := db.Init(dbUrl)

	tasks.RegisterRoutes(r,h)

	r.Run(port,)
}