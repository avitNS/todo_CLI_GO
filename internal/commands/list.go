package commands

import (
	"context"
	"fmt"
	"todo/internal/service"
)

type ListCommand struct {
	service service.Service
}

func NewListCommand(args []string, service service.Service) (service.Command, error) {
	return &ListCommand{service: service}, nil
}

func (cmd *ListCommand) Execute(ctx context.Context) error {

	tasks, err := cmd.service.List(ctx)

	if err != nil {
		return fmt.Errorf("command: error to load tasks list: %w", err)
	}

	for _, t := range tasks {
		flagDone := "| |"
		if t.Done {
			flagDone = "|x|"
		}
		fmt.Printf("%d - %v %v\n", t.ID, t.Title, flagDone)
	}

	return nil
}
