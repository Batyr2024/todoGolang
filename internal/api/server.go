package http

import (
	"github.com/Batyr2024/todoGolang/internal/api/handler"
	"github.com/Batyr2024/todoGolang/internal/api/middleware/corses"
	"github.com/Batyr2024/todoGolang/internal/api/middleware/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServer(h *handler.Task) *ServerHTTP {

	engine := gin.Default()

	engine.Use(gin.Recovery())
	engine.Use(corses.New())
	engine.Use(logger.New())

	staticPath, _ := filepath.Abs("/home/dunice/todo")
	fs := http.FileServer(http.Dir(staticPath))

	routes := engine.Group("/tasks")

	routes.GET("/", h.FindAll)
	routes.POST("/", h.Create)
	routes.DELETE("/:id", h.DeleteByID)
	routes.PATCH("/", h.ChangeCheckedByID)
	routes.PATCH("/:check", h.ChangeCheckedAll)
	routes.DELETE("/", h.DeleteAll)
	routes.PUT("/", h.ChangeText)
	routes.GET("/panic", h.Panicaaa)
	routes.GET("/ht", func() gin.HandlerFunc { return fs })
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	})

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &ServerHTTP{engine: engine}
}
func (s *ServerHTTP) Start(port string) {
	s.engine.Run(port)

}
