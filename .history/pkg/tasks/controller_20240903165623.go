package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct{
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *)