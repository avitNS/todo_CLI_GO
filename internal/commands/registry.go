package commands

import (
	"todo/internal/app"
)

type CommandFactory func([]string) (app.Command, error)

var Registry = map[string]CommandFactory{
	"add":    NewAddCommand,
	"done":   NewDoneCommand,
	"remove": NewRemoveCommand,
	"list":   NewListCommand,
}
