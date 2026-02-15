package parser

import (
	"os"
	"todo/internal/app"
	"todo/internal/commands"
)

func ParseArgs(args []string) (app.Command, error) {

	if len(args) == 0 {
		return nil, ErrUnknownCommand
	}

	if len(args) > 0 && args[0] == os.Args[0] {
		args = args[1:]
	}

	cmd, ok := commands.Registry[args[0]]

	if !ok {
		return nil, ErrUnknownCommand
	}

	return cmd(args[1:])

}
