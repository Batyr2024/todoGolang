package api

import (
	"github.com/Batyr2024/todoGolang/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var taskHandler = handlers.NewTaskHandler()

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &taskHandler{}

	routes := r.Group("/tasks")
		routes.GET("/", h.FindAll)
		routes.POST("/", h.Create)
		routes.DELETE("/",h.DeleteByID)
}
