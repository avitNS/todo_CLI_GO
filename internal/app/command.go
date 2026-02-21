package app

type Command interface {
	Execute(repo TaskRepository) error
}
