package task

import (
	"encoding/json"
	"os"
)

const filePath = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // file chưa có -> trả về rỗng
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
