package storage

import "todo/internal/model"

type TaskRepository interface {
	Save([]model.Task) error
	Load() ([]model.Task, error)
}
