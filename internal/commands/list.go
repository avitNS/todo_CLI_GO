package commands

import (
	"fmt"
	"todo/internal/storage"
)

type ListCommand struct{}

func (c *ListCommand) Execute(repo storage.TaskRepository) error {

	tasks, err := repo.Load()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks\n")
		return nil
	}

	for _, t := range tasks {

		done := "[]"
		if t.Done {
			done = "[x]"
		}
		fmt.Printf("%d %v %v\n", t.ID, done, t.Title)
	}

	return nil
}
