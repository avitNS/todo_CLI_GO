package commands

import (
	"todo/internal/service"
)

type CommandFactory func([]string, service.Service) (service.Command, error)

var Registry = map[string]CommandFactory{
	"add":    NewAddCommand,
	"done":   NewDoneCommand,
	"remove": NewRemoveCommand,
	"list":   NewListCommand,
}
