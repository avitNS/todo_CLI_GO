package service

import (
	"context"
	"todo/internal/model"
)

type TaskRepository interface {
	Save(context.Context, []model.Task) error
	Load(context.Context) ([]model.Task, error)
}
