package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Batyr2024/todoGolang/domain"
	"github.com/gin-gonic/gin"
)

type interfaceUseCase interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) error
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
}

type Task struct {
	useCase interfaceUseCase
}

func New(taskUseCase interfaceUseCase) *Task {
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
	queryId := c.Param("id")
	id, errParseInt := strconv.Atoi(queryId)
	if errParseInt != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}
	errRepo := h.useCase.DeleteByID(c.Request.Context(), id)
	if errRepo != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (h *Task) Create(c *gin.Context) {
	var dataTask domain.Task

	if err := c.BindJSON(&dataTask); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	err := h.useCase.Create(c.Request.Context(), dataTask)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (h *Task) ChangeCheckedByID(c *gin.Context) {
	queryId := c.Query("id")
	queryChecked := c.Query("check")
	id, errParseInt := strconv.Atoi(queryId)
	Checked, errParseBool := strconv.ParseBool(queryChecked)
	if errParseInt != nil && errParseBool != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id or checked",
		})
		return
	}
	errRep := h.useCase.ChangeCheckedByID(c.Request.Context(), id, Checked)
	if errRep != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (h *Task) ChangeCheckedAll(c *gin.Context) {
	queryChecked := c.Param("check")
	Checked, errParseBool := strconv.ParseBool(queryChecked)
	if errParseBool != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id or checked",
		})
		return
	}
	errRep := h.useCase.ChangeCheckedAll(c.Request.Context(), Checked)
	if errRep != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (h *Task) DeleteAll(c *gin.Context) {
	errRepo := h.useCase.DeleteAll(c.Request.Context())
	if errRepo != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, 1)
}
