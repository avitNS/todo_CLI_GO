package commands

import (
	"todo/internal/service"
)

type CommandFactory func([]string, *service.TaskService) (service.Command, error)

var Registry = map[string]CommandFactory{
	"add":    NewAddCommand,
	"done":   NewDoneCommand,
	"remove": NewRemoveCommand,
	"list":   NewListCommand,
}
