package parser

import (
	"fmt"
	"todo/internal/commands"
	"todo/internal/service"
)

func ParseArgs(args []string, service service.Service) (service.Command, error) {

	if len(args) == 0 {
		return nil, fmt.Errorf("parser: no command: %w", ErrMissingCommand)
	}

	cmd, ok := commands.Registry[args[0]]

	if !ok {
		return nil, fmt.Errorf("parser: no command: %w", ErrUnknownCommand)
	}

	return cmd(args[1:], service)

}
