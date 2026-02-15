package commands

import (
	"flag"
	"todo/internal/app"
	"todo/internal/storage"
)

type RemoveCommand struct {
	id int
}

func NewRemoveCommand(args []string) (app.Command, error) {
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

	return &RemoveCommand{id: id}, nil
}

func (cmd RemoveCommand) Execute(repo storage.TaskRepository) error {

	tasks, err := repo.Load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == cmd.id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err = repo.Save(tasks); err != nil {
				return err
			}
		}
	}

	return nil
}
