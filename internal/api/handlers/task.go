package handlers

import (
	"net/http"
	"strconv"

	"github.com/Batyr2024/todoGolang/db/models"
	services "github.com/Batyr2024/todoGolang/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct{
	taskUseCase services.TaskUseCase
}

func NewTaskHandler(usecase services.TaskUseCase) *TaskHandler{
	return &TaskHandler{taskUseCase: usecase}
}

func (cr *TaskHandler) FindAll(c *gin.Context){
	tasks,err := cr.taskUseCase.FindAll(c.Request.Context())
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,tasks)
}

func (cr *TaskHandler) DeleteByID(c *gin.Context){
	queryId := c.Query("id")
	id, err := strconv.Atoi(queryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}
	tasks,err := cr.taskUseCase.DeleteByID(c.Request.Context(),id)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,tasks)
}

func (cr *TaskHandler) Create(c *gin.Context){
	var dataTask models.Task

	if err:= c.BindJSON(&dataTask); err != nil{
		c.JSON(http.StatusNotFound,err)
		return
	} 
	tasks,err := cr.taskUseCase.Create(c.Request.Context(),dataTask)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,tasks)
} 