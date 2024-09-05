package usecase

import (
	"context"

	"github.com/Batyr2024/todoGolang/db/models"
	interfaces "github.com/Batyr2024/todoGolang/internal/repository/interfaces"
	services "github.com/Batyr2024/todoGolang/internal/usecase/interfaces"
)


type taskUseCase struct {
	taskRepo interfaces.TaskRepository
}

func NewTaskUseCase(repo interfaces.TaskRepository) services.TaskUseCase{
	return &taskUseCase{
		taskRepo: repo,
	}
}

func (c *taskUseCase) FindAll(ctx context.Context)([]models.Task,error){
	tasks,err := c.taskRepo.FindAll(ctx)
	return tasks,err
}

func (c *taskUseCase) Create(ctx context.Context,task models.Task)(bool,error){
	done,err := c.taskRepo.Create(ctx,task)
	return done,err
}

func (c *taskUseCase) DeleteByID(ctx context.Context, id int)(bool,error){
	done,err := c.taskRepo.DeleteByID(ctx,id)
	return done,err
}

