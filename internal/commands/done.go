package commands

import (
	"fmt"
	"todo/internal/model"
)

type DoneCommand struct {
	ID int
}

func (cmd *DoneCommand) Execute(tasks []model.Task) ([]model.Task, bool, error) {
	if cmd.ID <= 0 {
		return nil, false, fmt.Errorf("ID is incorrect")
	}

	for i, t := range tasks {
		if t.ID == cmd.ID {
			tasks[i].Done = true
			return tasks, true, nil
		}
	}
	return nil, false, nil
}
