package http

import (
	"github.com/Batyr2024/todoGolang/internal/api/handler"
	"github.com/Batyr2024/todoGolang/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServer(h *handler.Task) *ServerHTTP {
	engine := gin.Default()

	engine.Use(middleware.Cors())

	routes := engine.Group("/tasks")

	routes.GET("/", h.FindAll)
	routes.POST("/", h.Create)
	routes.DELETE("/:id", h.DeleteByID)
	routes.PATCH("/", h.ChangeCheckedByID)
	routes.PATCH("/:check", h.ChangeCheckedAll)
	routes.DELETE("/", h.DeleteAll)

	return &ServerHTTP{engine: engine}
}
func (s *ServerHTTP) Start(port string) {
	s.engine.Run(port)
}
