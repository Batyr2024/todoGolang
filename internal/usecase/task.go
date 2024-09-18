package usecase

import (
	"context"
	"github.com/Batyr2024/todoGolang/domain"
)

//go:generate mockgen -source=task.go -destination=../../mocks/TaskUsecase.go
type interfaceUseCase interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Create(ctx context.Context, task domain.Task) error
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
	ChangeText(ctx context.Context, id int, text string) error
}

type UseCase struct {
	taskRepo interfaceUseCase
}

func New(repo interfaceUseCase) interfaceUseCase {
	return &UseCase{
		taskRepo: repo,
	}
}

func (c *UseCase) FindAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := c.taskRepo.FindAll(ctx)
	return tasks, err
}

func (c *UseCase) Create(ctx context.Context, task domain.Task) error {
	err := c.taskRepo.Create(ctx, task)
	return err
}

func (c *UseCase) DeleteByID(ctx context.Context, id int) error {
	err := c.taskRepo.DeleteByID(ctx, id)
	return err
}
func (c *UseCase) ChangeCheckedByID(ctx context.Context, id int, checked bool) error {
	err := c.taskRepo.ChangeCheckedByID(ctx, id, checked)
	return err
}

func (c *UseCase) ChangeCheckedAll(ctx context.Context, checked bool) error {
	err := c.taskRepo.ChangeCheckedAll(ctx, checked)
	return err
}

func (c *UseCase) DeleteAll(ctx context.Context) error {
	err := c.taskRepo.DeleteAll(ctx)
	return err
}

func (c *UseCase) ChangeText(ctx context.Context, id int, text string) error {
	err := c.taskRepo.ChangeText(ctx, id, text)
	return err
}
