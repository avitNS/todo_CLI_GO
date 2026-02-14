package commands

import (
	"errors"
	"todo/internal/storage"
)

type RemoveCommand struct {
	ID int
}

func (cmd RemoveCommand) Execute(repo storage.TaskRepository) error {
	if cmd.ID <= 0 {
		return errors.New("ID is incorrect")
	}

	tasks, err := repo.Load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == cmd.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err = repo.Save(tasks); err != nil {
				return err
			}
		}
	}

	return nil
}
