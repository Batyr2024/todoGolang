package http

import (
	"github.com/Batyr2024/todoGolang/internal/api/handler"
	"github.com/Batyr2024/todoGolang/internal/api/middleware/corses"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServer(h *handler.Task) *ServerHTTP {

	engine := gin.Default()

	engine.Use(corses.New())
	engine.Use(gin.Recovery())

	routes := engine.Group("/tasks")

	routes.GET("/", h.FindAll)
	routes.POST("/", h.Create)
	routes.DELETE("/:id", h.DeleteByID)
	routes.PATCH("/", h.ChangeCheckedByID)
	routes.PATCH("/:check", h.ChangeCheckedAll)
	routes.DELETE("/", h.DeleteAll)
	routes.PUT("/", h.ChangeText)
	routes.GET("/panic", h.Panicaaa)

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &ServerHTTP{engine: engine}
}
func (s *ServerHTTP) Start(port string) {
	s.engine.Run(port)
}
