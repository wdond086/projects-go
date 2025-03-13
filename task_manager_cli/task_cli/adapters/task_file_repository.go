package adapters

import (
	"os"

	"github.com/wdond086/projects-go/task-manager-cli/task_cli/domain/task"
)

type TaskFileRepository struct {
	file os.File
}

func NewTaskFileRepository(file *os.File) TaskFileRepository {
	return TaskFileRepository{
		file: *file,
	}
}

func (repo TaskFileRepository) tasksCollection() []*task.Task {
	// Temporary imlementation
	tasks := []*task.Task{}
	return tasks
}
