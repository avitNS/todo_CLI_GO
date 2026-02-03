package commands

import (
	"fmt"
	"todo/internal/model"
)

type RemoveCommand struct {
	ID int
}

func (cmd RemoveCommand) Execute(tasks []model.Task) ([]model.Task, bool, error) {
	if cmd.ID <= 0 {
		return nil, false, fmt.Errorf("ID is incorrect")
	}

	for i, t := range tasks {
		if t.ID == cmd.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, true, nil
		}
	}

	return nil, false, nil
}
