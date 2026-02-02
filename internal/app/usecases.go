package app

import (
	"fmt"
	"time"
)

func listTask(tasks []Task) error {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	for _, t := range tasks {

		done := "[ ]"
		if t.Done {
			done = "[x]"
		}
		fmt.Printf("%s %d. %s\n", done, t.ID, t.Title)
	}
	return nil
}

func addTask(tasks []Task, title string) ([]Task, error) {
	if title == "" {
		return nil, fmt.Errorf("Title is necessary")
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	tasks = append(tasks, Task{
		ID:      maxID + 1,
		Title:   title,
		Done:    false,
		Created: time.Now(),
	})

	return tasks, nil
}

func removeTask(tasks []Task, id int) ([]Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID is incorrect")
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}

	return tasks, fmt.Errorf("ID not found")
}

func doneTask(tasks []Task, id int) ([]Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID is incorrect")
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			return tasks, nil
		}
	}

	return tasks, fmt.Errorf("ID not found")
}
