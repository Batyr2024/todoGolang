package interfaces

import (
	"context"

	"github.com/Batyr2024/todoGolang/db/models"
)



type TaskRepository interface{
	FindAll(ctx context.Context)([]models.Task,error)
	Create(ctx context.Context, task models.Task)(bool,error)
	DeleteByID(ctx context.Context, id int)(bool,error)

}