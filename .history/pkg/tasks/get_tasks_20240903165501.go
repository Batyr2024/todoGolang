package tasks

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Batyr2024/todoGolang/pkg/common/models"
)

func (h handler) GetTasks(c *gin.Context){
	var tasks []models.Task

	if result := h.DB.Find(&tasks); result.Error != nil{
		c.AbortWithError(http.StatusNotFound,result.Error)
	}

	c.JSON(http.Status)
}