package commands

import (
	"context"
	"flag"
	"todo/internal/service"
)

type RemoveCommand struct {
	service *service.TaskService
	id      int
}

func NewRemoveCommand(args []string, service *service.TaskService) (service.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingID
	}

	var id int
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	fs.IntVar(&id, "id", 0, "task id")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, ErrMissingID
	}

	return &RemoveCommand{id: id, service: service}, nil
}

func (cmd RemoveCommand) Execute(ctx context.Context) error {

	return cmd.service.Remove(ctx, cmd.id)
}
