package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	ErrUnknownCommand = errors.New("unknown command")
	ErrMissingTitle   = errors.New("missing title")
	ErrMissingID      = errors.New("missing ID")
)

// func declarFlags(args []string) (Flags, error) {
// 	fs := flag.NewFlagSet("todo", flag.ContinueOnError)

// 	var f Flags
// 	fs.StringVar(&f.Command, "cmd", "list", "command name string")
// 	fs.StringVar(&f.Title, "title", "", "task title string")
// 	fs.IntVar(&f.ID, "id", 0, "task id int")

// 	if err := fs.Parse(args); err != nil {
// 		return Flags{}, fmt.Errorf("Unknown arguments")
// 	}

// 	return f, nil

// }

func parseAdd(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingTitle
	}

	fmt.Println(args)

	var title string

	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.StringVar(&title, "title", "", "task title")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if title == "" {
		return nil, ErrMissingTitle
	}

	return &AddCommand{Title: title}, nil

}

func parseRemove(args []string) (Command, error) {
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

	return &RemoveCommand{ID: id}, nil

}

func parseDone(args []string) (Command, error) {
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

	return &DoneCommand{ID: id}, nil
}

// func parseArgs(args []string) (Command, error) {

// 	if len(args) > 0 && args[0] == os.Args[0] {
// 		args = args[1:]
// 	}

// 	f, err := declarFlags(args)
// 	if err != nil {
// 		return nil, err
// 	}

// 	switch f.Command {
// 	case "add":
// 		if len(f.Title) == 0 {
// 			return nil, ErrMissingTitle
// 		}

// 		return &AddCommand{
// 			Title: f.Title,
// 		}, nil

// 	case "done":
// 		if f.ID == 0 {
// 			return nil, ErrMissingID
// 		}

// 		return &DoneCommand{
// 			ID: f.ID,
// 		}, nil

// 	case "remove":
// 		if f.ID == 0 {
// 			return nil, ErrMissingID
// 		}

// 		return &RemoveCommand{
// 			ID: f.ID,
// 		}, nil

// 	case "list":
// 		return &ListCommand{}, nil

// 	default:
// 		return nil, ErrUnknownCommand
// 	}

// }

func parseArgs(args []string) (Command, error) {

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
		return &ListCommand{}, nil

	default:
		return nil, ErrUnknownCommand
	}

}
