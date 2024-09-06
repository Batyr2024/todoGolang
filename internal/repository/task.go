package repository

import (
	"context"

	"github.com/Batyr2024/todoGolang/domain"
	"gorm.io/gorm"
)

type interfaceHandler interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) (bool, error)
	DeleteByID(ctx context.Context, id int) (bool, error)
}

type dataBase struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) interfaceHandler {
	return &dataBase{DB}
}

func (c *dataBase) FindAll(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	err := c.DB.Find(&tasks).Error
	return tasks, err
}

func (c *dataBase) Create(ctx context.Context, task domain.Task) (bool, error) {
	err := c.DB.Create(&task).Error
	return true, err
}

func (c *dataBase) DeleteByID(ctx context.Context, id int) (bool, error) {
	var tasks domain.Task
	err := c.DB.Delete(&tasks, id).Error
	return true, err
}
