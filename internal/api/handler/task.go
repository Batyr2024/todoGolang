package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Batyr2024/todoGolang/domain"
	"github.com/gin-gonic/gin"
)

type interfaceHandler interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) (bool, error)
	DeleteByID(ctx context.Context, id int) (bool, error)
}

type Task struct {
	useCase interfaceHandler
}

func New(taskUseCase interfaceHandler) *Task {
	return &Task{useCase: taskUseCase}
}

func (h *Task) FindAll(c *gin.Context) {
	tasks, err := h.useCase.FindAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Task) DeleteByID(c *gin.Context) {
	queryId := c.Query("id")
	id, err := strconv.Atoi(queryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}
	tasks, err := h.useCase.DeleteByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Task) Create(c *gin.Context) {
	var dataTask domain.Task

	if err := c.BindJSON(&dataTask); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	tasks, err := h.useCase.Create(c.Request.Context(), dataTask)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, tasks)
}
