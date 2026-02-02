package storage

import (
	"encoding/json"
	"os"
)

func loadTasks() ([]Task, error) {

	f, err := os.ReadFile(JsonPath)
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

func saveTasks(tasks []Task) error {

	buf, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}

	tmp := JsonPath + ".tmp"
	if err := os.WriteFile(tmp, buf, 0644); err != nil {
		return err
	}

	if err := os.Rename(tmp, JsonPath); err != nil {
		return err
	}

	return nil
}
