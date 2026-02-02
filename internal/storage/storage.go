package storage

type Storage interface {
	Add(tasks []Task) error
	List() ([]Task, error)
}
