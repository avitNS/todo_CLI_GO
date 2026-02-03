package commands

import (
	"fmt"
	"todo/internal/model"
)

type ListCommand struct{}

func (c *ListCommand) Execute(tasks []model.Task) ([]model.Task, bool, error) {
	if len(tasks) == 0 {
		return nil, false, fmt.Errorf("No tasks")
	}

	for _, t := range tasks {

		done := "[]"
		if t.Done {
			done = "[x]"
		}
		fmt.Printf("%s %d. %s\n", done, t.ID, t.Title)
	}
	return nil, false, nil
}
