package usecase

import (
	"context"
	"github.com/Batyr2024/todoGolang/domain"
)

type interfaceHandler interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Create(ctx context.Context, task domain.Task) error
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
	ChangeText(ctx context.Context, id int, text string) error
}

type useCase struct {
	taskRepo interfaceHandler
}

func New(repo interfaceHandler) interfaceHandler {
	return &useCase{
		taskRepo: repo,
	}
}

func (c *useCase) FindAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := c.taskRepo.FindAll(ctx)
	return tasks, err
}

func (c *useCase) Create(ctx context.Context, task domain.Task) error {
	err := c.taskRepo.Create(ctx, task)
	return err
}

func (c *useCase) DeleteByID(ctx context.Context, id int) error {
	err := c.taskRepo.DeleteByID(ctx, id)
	return err
}
func (c *useCase) ChangeCheckedByID(ctx context.Context, id int, checked bool) error {
	err := c.taskRepo.ChangeCheckedByID(ctx, id, checked)
	return err
}

func (c *useCase) ChangeCheckedAll(ctx context.Context, checked bool) error {
	err := c.taskRepo.ChangeCheckedAll(ctx, checked)
	return err
}

func (c *useCase) DeleteAll(ctx context.Context) error {
	err := c.taskRepo.DeleteAll(ctx)
	return err
}

func (c *useCase) ChangeText(ctx context.Context, id int, text string) error {
	err := c.taskRepo.ChangeText(ctx, id, text)
	return err
}
