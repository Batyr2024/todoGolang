package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/rs/cors"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/tasks")
	routes.GET("/get", h.GetTasks)
	routes.POST("/post", h.AddTask)
}
