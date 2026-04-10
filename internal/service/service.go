package service

import (
	"context"
	"todo/internal/model"
)

type Service interface {
	Add(ctx context.Context, title string) error
	Done(ctx context.Context, id int) error
	Remove(ctx context.Context, id int) error
	List(ctx context.Context) ([]model.Task, error)
}
