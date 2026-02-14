package app

import (
	"todo/internal/storage"
)

type Command interface {
	Execute(repo storage.TaskRepository) error
}
