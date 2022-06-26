package models

import "time"

type Tasks struct {
	Tasks []Task
}

func (task *Task) GetTasks() ([]*Task, error) {
	var tasks []*Task
	// Note: This is to mimic delay
	time.Sleep(time.Second * 1 / 2)
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, nTask := range tasks {
		err := nTask.setTaskByIDToCache()
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}