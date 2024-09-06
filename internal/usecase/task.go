package usecase

import (
	"context"
	"github.com/Batyr2024/todoGolang/domain"
)

type interfaceHandler interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) (bool, error)
	DeleteByID(ctx context.Context, id int) (bool, error)
}

type useCase struct {
	taskRepo interfaceHandler
}

func New(repo interfaceHandler) interfaceHandler {
	return &useCase{
		taskRepo: repo,
	}
}

func (c *useCase) FindAll(ctx context.Context) ([]domain.Task, error) {
	tasks, err := c.taskRepo.FindAll(ctx)
	return tasks, err
}

func (c *useCase) Create(ctx context.Context, task domain.Task) (bool, error) {
	done, err := c.taskRepo.Create(ctx, task)
	return done, err
}

func (c *useCase) DeleteByID(ctx context.Context, id int) (bool, error) {
	done, err := c.taskRepo.DeleteByID(ctx, id)
	return done, err
}
