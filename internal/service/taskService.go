package service

import (
	"context"
	"todo/internal/model"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Add(ctx context.Context, title string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if title == "" {
		return ErrMissingTitle
	}

	tasks, err := s.repo.Load(ctx)
	if err != nil {
		return err
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:    newID,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, newTask)

	return s.repo.Save(ctx, tasks)
}

func (s *TaskService) Done(ctx context.Context, id int) error
func (s *TaskService) Remove(ctx context.Context, id int) error
func (s *TaskService) List(ctx context.Context) error
