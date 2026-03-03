package commands

import (
	"context"
	"flag"
	"todo/internal/service"
)

type AddCommand struct {
	service *service.TaskService
	title   string
}

func NewAddCommand(args []string, service *service.TaskService) (service.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingTitle
	}

	var title string

	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.StringVar(&title, "title", "", "task title")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if title == "" {
		return nil, ErrMissingTitle
	}

	return &AddCommand{title: title, service: service}, nil
}

func (cmd *AddCommand) Execute(ctx context.Context) error {

	return cmd.service.Add(ctx, cmd.title)

}
