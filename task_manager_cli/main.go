package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wdond086/projects-go/task-manager-cli/task_cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = task_cli.WithLogger(ctx, task_cli.NewLogger())
	ctx = task_cli.WithTransactionId(ctx)
	logger := task_cli.FromContext(ctx)
	logger.Info(fmt.Sprintln(os.Args))
}
