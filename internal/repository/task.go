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
	Create(ctx context.Context, task domain.Task) (*domain.Task, error)
	DeleteByID(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ChangeCheckedByID(ctx context.Context, id int, checked bool) error
	ChangeCheckedAll(ctx context.Context, checked bool) error
	ChangeText(ctx context.Context, id int, text string) error
}
type Repository struct {
	db *DB.Queries
	interfaceRepository
}

func NewDBTX(conn DB.DBTX) interfaceRepository {
	return &Repository{db: DB.New(conn)}
}
func New(conn *pgx.Conn) interfaceRepository {
	return &Repository{db: DB.New(conn)}
}

func (r *Repository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := r.db.GetAll(ctx)
	return tasks, err
}

func (r *Repository) Create(ctx context.Context, task domain.Task) (*domain.Task, error) {
	newTask, err := r.db.Create(ctx, &DB.CreateParams{Text: task.Text})
	return newTask, err
}

func (r *Repository) DeleteByID(ctx context.Context, ID int) error {
	_, err := r.db.DeleteByID(ctx, &DB.DeleteByIDParams{ID: int32(ID)})
	return err
}

func (r *Repository) ChangeCheckedByID(ctx context.Context, ID int, Checked bool) error {
	_, err := r.db.ChangeCheckedByID(ctx, &DB.ChangeCheckedByIDParams{ID: int32(ID), Checked: &Checked})
	return err
}

func (r *Repository) ChangeCheckedAll(ctx context.Context, Checked bool) error {
	_, err := r.db.ChangeCheckedAll(ctx, &DB.ChangeCheckedAllParams{Checked: &Checked})
	return err
}

func (r *Repository) DeleteAll(ctx context.Context) error {
	_, err := r.db.DeleteAll(ctx)
	return err
}

func (r *Repository) ChangeText(ctx context.Context, ID int, Text string) error {
	_, err := r.db.ChangeTextByID(ctx, &DB.ChangeTextByIDParams{ID: int32(ID), Text: Text})
	return err
}
