package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Batyr2024/todoGolang/domain"
	"github.com/gin-gonic/gin"
)

type interfaceUseCase interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
	ChangeText(ctx context.Context, id int, text string) error
}

type interfaceHandler interface {
	FindAll(ctx context.Context)
	Create(ctx context.Context)
	DeleteByID(ctx context.Context)
	DeleteAll(ctx context.Context)
	ChangeCheckedByID(ctx context.Context)
	ChangeCheckedAll(ctx context.Context)
	ChangeText(ctx context.Context)
}

type Task struct {
	useCase interfaceUseCase
	interfaceHandler
}

func New(taskUseCase interfaceUseCase) *Task {
	return &Task{useCase: taskUseCase}
}

// FindAll            godoc
// @Summary      Get tasks array
// @Description  Responds with the list of all tasks as JSON.
// @Tags         tasks
// @Produce      json
// @Success      200  {array}  domain.Task
// @Failure 	 400 {string} string "error"
// @Router       /tasks [get]
func (h *Task) FindAll(c *gin.Context) {
	tasks, err := h.useCase.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// DeleteByID            godoc
// @Summary      Delete single task by id
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      204  {int}  1
// @Failure 	 400,404,500 {string} string "error"
// @Router       /tasks/{id} [delete]
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
		c.JSON(http.StatusNotFound, errRepo)
		return
	}
	c.JSON(http.StatusNoContent, 1)
}

// Create            godoc
// @Summary      Create task
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      201  {int}  1
// @Failure 	 400 {string} string "error"
// @Router       /tasks [post]
func (h *Task) Create(c *gin.Context) {
	var dataTask domain.Task
	if err := c.BindJSON(&dataTask); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if dataTask.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text field is missing"})
		return
	}
	newTask, err := h.useCase.Create(c.Request.Context(), dataTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// ChangeCheckedByID            godoc
// @Summary      Change checked task by ID
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      204  {int}  1
// @Failure 	 400,404,500 {string} string "error"
// @Router       /tasks [patch]
func (h *Task) ChangeCheckedByID(c *gin.Context) {
	queryId := c.Query("id")
	queryChecked := c.Query("check")
	if queryId == "" || queryChecked == "" {
		c.JSON(http.StatusBadRequest, `The "Checked" or "Id" field is missing`)
		return
	}
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
		c.JSON(http.StatusNotFound, errRep)
		return
	}
	c.JSON(http.StatusNoContent, 1)
}

// ChangeCheckedAll            godoc
// @Summary      Change checked all task
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      204  {int}  1
// @Failure 	 400,404,500 {string} string "error"
// @Router       /tasks/{checked} [patch]
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
	c.JSON(http.StatusNoContent, 1)
}

// DeleteAll            godoc
// @Summary      Delete all checked true tasks
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      204  {int}  1
// @Failure 	 404 {string} string "error"
// @Router       /tasks [delete]
func (h *Task) DeleteAll(c *gin.Context) {
	errRepo := h.useCase.DeleteAll(c.Request.Context())
	if errRepo != nil {
		c.JSON(http.StatusNotFound, errRepo)
		return
	}
	c.JSON(http.StatusNoContent, 1)
}

// ChangeText           godoc
// @Summary      Change text task
// @Description  Responds 1 in JSON format.
// @Tags         tasks
// @Produce      json
// @Success      200  {int}  1
// @Failure 	 400,404,500 {string} string "error"
// @Router       /tasks [put]
func (h *Task) ChangeText(c *gin.Context) {
	var dataTask domain.Task
	if err := c.BindJSON(&dataTask); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if dataTask.ID == 0 || dataTask.Text == "" {
		c.JSON(http.StatusBadRequest, `The "Text" or "Id" field is missing`)
		return
	}
	errRepo := h.useCase.ChangeText(c.Request.Context(), int(dataTask.ID), dataTask.Text)
	if errRepo != nil {
		c.JSON(http.StatusNotFound, errRepo)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func (h *Task) Panicaaa(c *gin.Context) {
	panic("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
}
