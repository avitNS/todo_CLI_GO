package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"todo/internal/model"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

func (file *FileStorage) Load(ctx context.Context) ([]model.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	f, err := os.ReadFile(file.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, fmt.Errorf("repo: failed to read file: %w", err)
	}

	var tasks []model.Task
	err = json.Unmarshal(f, &tasks)
	if err != nil {
		return nil, fmt.Errorf("repo: failed to unmarshal: %w", err)
	}
	return tasks, nil
}

func (file *FileStorage) Save(ctx context.Context, tasks []model.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	buf, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return fmt.Errorf("repo: failed to marshal: %w", err)
	}

	tmp := file.path + ".tmp"
	if err := os.WriteFile(tmp, buf, 0644); err != nil {
		return fmt.Errorf("repo: failed to write to file: %w", err)
	}

	if err := os.Rename(tmp, file.path); err != nil {
		return fmt.Errorf("repo: failed to rename: %w", err)
	}

	return nil
}
