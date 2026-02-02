package main

type Command interface {
	Execute(tasks []Task) ([]Task, bool, error)
}

type ListCommand struct{}

func (c ListCommand) Execute(tasks []Task) ([]Task, bool, error) {
	return nil, false, listTask(tasks)
}

type AddCommand struct {
	Title string
}

func (c AddCommand) Execute(tasks []Task) ([]Task, bool, error) {
	tasks, err := addTask(tasks, c.Title)
	return tasks, true, err
}

type DoneCommand struct {
	ID int
}

func (c DoneCommand) Execute(tasks []Task) ([]Task, bool, error) {
	tasks, err := doneTask(tasks, c.ID)
	return tasks, true, err
}

type RemoveCommand struct {
	ID int
}

func (c RemoveCommand) Execute(tasks []Task) ([]Task, bool, error) {
	tasks, err := removeTask(tasks, c.ID)
	return tasks, true, err
}
