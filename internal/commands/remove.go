package commands

import (
	"context"
	"flag"
	"fmt"
	"todo/internal/service"
)

type RemoveCommand struct {
	service *service.TaskService
	id      int
}

func NewRemoveCommand(args []string, service *service.TaskService) (service.Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	var id int
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	fs.IntVar(&id, "id", 0, "task id")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("command: failed to parse command: %w", err)
	}

	if id <= 0 {
		return nil, fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	return &RemoveCommand{id: id, service: service}, nil
}

func (cmd RemoveCommand) Execute(ctx context.Context) error {

	return cmd.service.Remove(ctx, cmd.id)
}
