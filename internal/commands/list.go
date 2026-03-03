package commands

import (
	"context"
	"todo/internal/service"
)

type ListCommand struct {
	service *service.TaskService
}

func NewListCommand(args []string, service *service.TaskService) (service.Command, error) {
	return &ListCommand{service: service}, nil
}

func (cmd *ListCommand) Execute(ctx context.Context) error {

	return cmd.service.List(ctx)
}
