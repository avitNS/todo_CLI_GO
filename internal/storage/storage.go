package storage

import "todo/internal/model"

type Storage interface {
	Add(tasks []model.Task) error
	List() ([]model.Task, error)
}
