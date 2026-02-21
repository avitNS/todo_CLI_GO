package commands

import (
	"flag"
	"todo/internal/app"
)

type DoneCommand struct {
	id int
}

func NewDoneCommand(args []string) (app.Command, error) {
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

	return &DoneCommand{id: id}, nil
}

func (cmd *DoneCommand) Execute(repo app.TaskRepository) error {

	tasks, err := repo.Load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == cmd.id {
			tasks[i].Done = true
			if err = repo.Save(tasks); err != nil {
				return err
			}
		}
	}
	return nil
}
