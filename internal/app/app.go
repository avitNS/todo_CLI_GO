package app

import (
	"todo/internal/storage"
)

type App struct {
	repo storage.TaskRepository
}

func NewApp(repo storage.TaskRepository) *App {
	return &App{repo: repo}
}

func (app *App) Execute(cmd Command) error {

	err := cmd.Execute(app.repo)
	if err != nil {
		return err
	}

	return nil
}
