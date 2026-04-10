package commands

import (
	"context"
	"flag"
	"fmt"
	"todo/internal/service"
)

type RemoveCommand struct {
	service service.Service
	id      int
}

func NewRemoveCommand(args []string, service service.Service) (service.Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("command: args are empty: %w", ErrMissingID)
	}

	var id int
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	fs.IntVar(&id, "id", 0, "task id")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("command: failed to parse command: %w", err)
	}

	if id <= 0 {
		return nil, fmt.Errorf("command: failed to remove task: %w", ErrInvalidID)
	}

	return &RemoveCommand{id: id, service: service}, nil
}

func (cmd *RemoveCommand) Execute(ctx context.Context) error {

	return cmd.service.Remove(ctx, cmd.id)
}
