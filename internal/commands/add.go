package commands

import (
	"errors"
	"fmt"
	"time"
	"todo/internal/model"
	"todo/internal/storage"
)

type AddCommand struct {
	Title string
}

func (cmd *AddCommand) Execute(repo storage.TaskRepository) error {

	if cmd.Title == "" {
		fmt.Printf("Title is missing\n")
		return errors.New("Title is missing")
	}

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
		Title:   cmd.Title,
		Done:    false,
		Created: time.Now(),
	})

	if err = repo.Save(tasks); err != nil {
		return err
	}

	return nil

}
