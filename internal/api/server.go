package api

import (
	"github.com/Batyr2024/todoGolang/internal/api/handlers"
	"github.com/Batyr2024/todoGolang/internal/repository"
	"github.com/Batyr2024/todoGolang/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewTaskHandler(usecase.NewTaskUseCase(repository.NewTaskRepository(db)))

	routes := r.Group("/tasks")
		routes.GET("/", h.FindAll)
		routes.POST("/", h.Create)
		routes.DELETE("/",h.DeleteByID)
}
