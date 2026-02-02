package storage

import (
	"encoding/json"
	"os"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

func (file *FileStorage) loadTasks() ([]Task, error) {

	f, err := os.ReadFile(file.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(f, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (file *FileStorage) saveTasks(tasks []Task) error {

	buf, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}

	tmp := file.path + ".tmp"
	if err := os.WriteFile(tmp, buf, 0644); err != nil {
		return err
	}

	if err := os.Rename(tmp, file.path); err != nil {
		return err
	}

	return nil
}
