package commands

import (
	"context"
	"flag"
	"todo/internal/service"
)

type DoneCommand struct {
	service *service.TaskService
	id      int
}

func NewDoneCommand(args []string, service *service.TaskService) (service.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingID
	}

	var id int
	fs := flag.NewFlagSet("done", flag.ContinueOnError)
	fs.IntVar(&id, "id", 0, "task id")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, ErrMissingID
	}

	return &DoneCommand{id: id, service: service}, nil
}

func (cmd *DoneCommand) Execute(ctx context.Context) error {

	return cmd.service.Done(ctx, cmd.id)
}
