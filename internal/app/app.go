package app

import (
	"todo/internal/storage"
)

type App struct {
	storage storage.Storage
}

func NewApp(storage storage.Storage) *App {
	return &App{storage: storage}
}

func (app *App) Execute(cmd Command) error {
	tasks, err := app.storage.List()

	if err != nil {
		return err
	}

	tasks, mutated, err := cmd.Execute(tasks)
	if err != nil {
		return err
	}

	if mutated {
		return app.storage.Add(tasks)
	}

	return nil
}
