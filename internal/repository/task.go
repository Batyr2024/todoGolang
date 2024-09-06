package repository

import (
	"context"

	"github.com/Batyr2024/todoGolang/domain"
	"gorm.io/gorm"
)

type interfaceHandler interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) error
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
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

func (c *dataBase) Create(ctx context.Context, task domain.Task) error {
	err := c.DB.Create(&task).Error
	return err
}

func (c *dataBase) DeleteByID(ctx context.Context, id int) error {
	var tasks domain.Task
	err := c.DB.Delete(&tasks, id).Error
	return err
}

func (c *dataBase) ChangeCheckedByID(ctx context.Context, id int, checked bool) error {
	var tasks domain.Task
	err := c.DB.Model(&tasks).Where("id = ?", id).Update("checked", checked).Error
	return err
}

func (c *dataBase) ChangeCheckedAll(ctx context.Context, checked bool) error {
	var tasks domain.Task
	err := c.DB.Model(&tasks).Where("0=?", 0).Update("checked", checked).Error
	return err
}

func (c *dataBase) DeleteAll(ctx context.Context) error {
	var tasks domain.Task
	err := c.DB.Model(&tasks).Where("checked=?", true).Delete(&tasks).Error
	return err
}
