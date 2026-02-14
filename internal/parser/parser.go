package parser

import (
	"errors"
	"flag"
	"os"
	"todo/internal/app"
	"todo/internal/commands"
)

var (
	ErrUnknownCommand = errors.New("unknown command")
	ErrMissingTitle   = errors.New("missing title")
	ErrMissingID      = errors.New("missing ID")
)

func ParseArgs(args []string) (app.Command, error) {

	if len(args) == 0 {
		return nil, ErrUnknownCommand
	}

	if len(args) > 0 && args[0] == os.Args[0] {
		args = args[1:]
	}

	switch args[0] {
	case "add":
		return parseAdd(args[1:])

	case "done":
		return parseDone(args[1:])

	case "remove":
		return parseRemove(args[1:])

	case "list":
		return &commands.ListCommand{}, nil

	default:
		return nil, ErrUnknownCommand
	}

}

func parseAdd(args []string) (app.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingTitle
	}

	// fmt.Println(args)

	var title string

	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.StringVar(&title, "title", "", "task title")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if title == "" {
		return nil, ErrMissingTitle
	}

	return &commands.AddCommand{Title: title}, nil

}

func parseRemove(args []string) (app.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingID
	}

	var id int
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	fs.IntVar(&id, "id", 0, "task id")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, ErrMissingID
	}

	return &commands.RemoveCommand{ID: id}, nil

}

func parseDone(args []string) (app.Command, error) {
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

	return &commands.DoneCommand{ID: id}, nil
}
