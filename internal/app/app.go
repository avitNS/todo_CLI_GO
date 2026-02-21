package app

type App struct {
	repo TaskRepository
}

func NewApp(repo TaskRepository) *App {
	return &App{repo: repo}
}

func (app *App) Execute(cmd Command) error {

	err := cmd.Execute(app.repo)
	if err != nil {
		return err
	}

	return nil
}
