package repository

import (
	"context"
	DB "github.com/Batyr2024/todoGolang/db/sqlc"
	"github.com/jackc/pgx/v5"

	"github.com/Batyr2024/todoGolang/domain"
)

//go:generate mockgen -source=task.go -destination=../../mocks/TaskRepository.go

type interfaceRepository interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Create(ctx context.Context, task domain.Task) error
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
	ChangeText(ctx context.Context, id int, text string) error
}
type repository struct {
	db *DB.Queries
}

func New(conn *pgx.Conn) interfaceRepository {
	return &repository{DB.New(conn)}
}

func (r *repository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := r.db.GetAll(ctx)
	return tasks, err
}

func (r *repository) Create(ctx context.Context, task domain.Task) error {
	_, err := r.db.Create(ctx, &DB.CreateParams{Text: task.Text})
	return err
}

func (r *repository) DeleteByID(ctx context.Context, ID int) error {
	_, err := r.db.DeleteByID(ctx, &DB.DeleteByIDParams{ID: int32(ID)})
	return err
}

func (r *repository) ChangeCheckedByID(ctx context.Context, ID int, Checked bool) error {
	_, err := r.db.ChangeCheckedByID(ctx, &DB.ChangeCheckedByIDParams{ID: int32(ID), Checked: &Checked})
	return err
}

func (r *repository) ChangeCheckedAll(ctx context.Context, Checked bool) error {
	_, err := r.db.ChangeCheckedAll(ctx, &DB.ChangeCheckedAllParams{Checked: &Checked})
	return err
}

func (r *repository) DeleteAll(ctx context.Context) error {
	_, err := r.db.DeleteAll(ctx)
	return err
}

func (r *repository) ChangeText(ctx context.Context, ID int, Text string) error {
	_, err := r.db.ChangeTextByID(ctx, &DB.ChangeTextByIDParams{ID: int32(ID), Text: Text})
	return err
}
