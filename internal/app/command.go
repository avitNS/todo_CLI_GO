package app

import "todo/internal/model"

type Command interface {
	Execute(tasks []model.Task) ([]model.Task, bool, error)
}
