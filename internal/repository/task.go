package repository

import (
	"context"

	"github.com/Batyr2024/todoGolang/db/models"
	"github.com/Batyr2024/todoGolang/internal/repository/interfaces"
	"gorm.io/gorm"
)

type taskDataBase struct {
	DB *gorm.DB
}

func NewTaskRepository(DB *gorm.DB) interfaces.TaskRepository{
	return &taskDataBase{DB}
}

func (c *taskDataBase) FindAll(ctx context.Context)([]models.Task,error){
	var tasks []models.Task
	err := c.DB.Find(&tasks).Error
	return tasks,err
}

func (c *taskDataBase) Create(ctx context.Context,task models.Task)(bool,error){
	err := c.DB.Create(&task).Error
	return true,err
}

func (c *taskDataBase) DeleteByID(ctx context.Context, id int)(bool,error){
	var tasks models.Task
	err := c.DB.Delete(&tasks,id).Error
	return true,err
}

