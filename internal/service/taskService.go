package service

import (
	"context"
	"fmt"
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
		return fmt.Errorf("command: failed to add task: %w", ErrMissingTitle)
	}

	tasks, err := s.repo.Load(ctx)
	if err != nil {
		return fmt.Errorf("command: failed to load tasks: %w", err)
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

func (s *TaskService) Done(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if id <= 0 {
		return fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	tasks, err := s.repo.Load(ctx)
	if err != nil {
		return fmt.Errorf("command: failed to load tasks: %w", err)
	}

	foundIdx := -1
	for i, t := range tasks {
		if t.ID == id {
			foundIdx = i
			break
		}
	}

	if foundIdx == -1 {
		return fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	tasks[foundIdx].Done = true

	return s.repo.Save(ctx, tasks)

}

func (s *TaskService) Remove(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if id <= 0 {
		return fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	tasks, err := s.repo.Load(ctx)
	if err != nil {
		return fmt.Errorf("command: failed to load tasks: %w", err)
	}

	foundIdx := -1
	for i, t := range tasks {
		if t.ID == id {
			foundIdx = i
			break
		}
	}

	if foundIdx == -1 {
		return fmt.Errorf("command: failed to add task: %w", ErrMissingID)
	}

	tasks = append(tasks[:foundIdx], tasks[foundIdx+1:]...)

	return s.repo.Save(ctx, tasks)
}

func (s *TaskService) List(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	tasks, err := s.repo.Load(ctx)
	if err != nil {
		return fmt.Errorf("command: failed to load tasks: %w", err)
	}

	for _, t := range tasks {
		flagDone := "| |"
		if t.Done {
			flagDone = "|x|"
		}
		fmt.Printf("%d - %v %v\n", t.ID, t.Title, flagDone)
	}

	return nil

}
