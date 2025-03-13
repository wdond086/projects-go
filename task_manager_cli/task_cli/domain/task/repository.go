package task

import "context"

type Repository interface {
	GetTask(ctx context.Context, id string) (*Task, error)
	GetTasks(ctx context.Context, status Status) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, id string) error
}
