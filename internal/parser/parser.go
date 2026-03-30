package parser

import (
	"fmt"
	"os"
	"todo/internal/commands"
	"todo/internal/service"
)

func ParseArgs(args []string, service *service.TaskService) (service.Command, error) {

	if len(args) == 0 {
		return nil, fmt.Errorf("parser: no command: %w", ErrUnknownCommand)
	}

	if len(args) > 0 && args[0] == os.Args[0] {
		args = args[1:]
	}

	cmd, ok := commands.Registry[args[0]]

	if !ok {
		return nil, fmt.Errorf("parser: no command: %w", ErrUnknownCommand)
	}

	return cmd(args[1:], service)

}
