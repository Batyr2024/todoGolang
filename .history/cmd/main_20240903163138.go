package main

import (
	"github.com/Batyr2024/todoGolang/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	
	r := gin.Default()
	db.Init(dbUrl)

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"port": port,
			"dbUrl": dbUrl,
		})
	})

	r.Run(port)
}