package services

import (
	"gin-gonic-tut/models"
	"github.com/google/uuid"
)

var (
	TaskService taskServiceInterface = &taskService{}
)

type taskServiceInterface interface {
	GetTasks() ([]*models.Task, error)
	GetTask(string) (*models.Task, error)
	CreateTask(models.Task) (*models.Task, error)
}

type taskService struct{}

func (t taskService) GetTask(id string) (*models.Task, error) {
	taskId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	task := &models.Task{
		Id: taskId,
	}
	task.Get()
	return task, nil
}

func (t taskService) CreateTask(task models.Task) (*models.Task, error) {
	task.Id = uuid.New()
	task.Create()
	return &task, nil
}

func (t taskService) GetTasks() ([]*models.Task, error) {
	var task models.Task
	return task.GetTasks()
}