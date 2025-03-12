package task

import "context"

type Repository interface {
	GetTask(ctx context.Context) (*Task, error)
	GetTasks(ctx context.Context) ([]*Task, error)
	CreateTask(ctx context.Context) (*Task, error)
	UpdateTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
}
