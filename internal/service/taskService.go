package service

import "context"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Add(ctx context.Context, title string) error
func (s *TaskService) Done(ctx context.Context, id int) error
func (s *TaskService) Remove(ctx context.Context, id int) error
func (s *TaskService) List(ctx context.Context) error
