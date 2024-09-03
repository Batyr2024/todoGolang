package tasks

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Batyr2024/todoGolang/pkg/common/models"
)

func (h *handler) DeleteTask(c *gin.Context){
	queryId, havingId := c.GetQuery("id")
	if havingId!=true&&queryId=="" {
		c.AbortWithStatus(404)
		return
	}

	if result := h.DB.Delete(&models.Task{},queryId); result.Error != nil{
		c.AbortWithError(http.StatusNotFound,result.Error)
		return
	}
}