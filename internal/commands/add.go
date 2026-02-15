package commands

import (
	"flag"
	"time"
	"todo/internal/app"
	"todo/internal/model"
	"todo/internal/storage"
)

type AddCommand struct {
	title string
}

func NewAddCommand(args []string) (app.Command, error) {
	if len(args) == 0 {
		return nil, ErrMissingTitle
	}

	var title string

	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.StringVar(&title, "title", "", "task title")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if title == "" {
		return nil, ErrMissingTitle
	}

	return &AddCommand{title: title}, nil
}

func (cmd *AddCommand) Execute(repo storage.TaskRepository) error {

	tasks, err := repo.Load()

	if err != nil {
		return err
	}
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	tasks = append(tasks, model.Task{
		ID:      maxID + 1,
		Title:   cmd.title,
		Done:    false,
		Created: time.Now(),
	})

	if err = repo.Save(tasks); err != nil {
		return err
	}

	return nil

}
