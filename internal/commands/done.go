package commands

import (
	"errors"
	"todo/internal/storage"
)

type DoneCommand struct {
	ID int
}

func (cmd *DoneCommand) Execute(repo storage.TaskRepository) error {
	if cmd.ID <= 0 {
		return errors.New("ID is incorrect")
	}

	tasks, err := repo.Load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == cmd.ID {
			tasks[i].Done = true
			if err = repo.Save(tasks); err != nil {
				return err
			}
		}
	}
	return nil
}
