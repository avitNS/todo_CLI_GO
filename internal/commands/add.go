package commands

import (
	"fmt"
	"time"
	"todo/internal/model"
)

type AddCommand struct {
	Title string
}

func (cmd *AddCommand) Execute(tasks []model.Task) ([]model.Task, bool, error) {

	if cmd.Title == "" {
		return nil, false, fmt.Errorf("Title is necessary")
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	tasks = append(tasks, model.Task{
		ID:      maxID + 1,
		Title:   cmd.Title,
		Done:    false,
		Created: time.Now(),
	})

	return tasks, true, nil
}
