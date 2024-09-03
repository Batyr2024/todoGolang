package tasks

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Batyr2024/todoGolang/pkg/common/models"
)

type AddTaskRequestBody struct{
	Text string `json:"text"`
}

func (h handler) AddTask(c *gin.Context){
	body :=AddTaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Text = body.Text

	if result := h.DB.Create(&task); result.Error != nil{
		c.AbortWithError(http.StatusNotFound,result.Error)
		return
	}
	
}