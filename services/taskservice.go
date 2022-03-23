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
	GetTaskByID(id uuid.UUID) (*models.Task, error)
	CreateTask(models.Task) (*models.Task, error)
	DeleteTask(task models.Task) error
	UpdateTask(task models.Task) (*models.Task, error)
}

type taskService struct{}

func (t taskService) GetTaskByID(id uuid.UUID) (*models.Task, error) {
	task := &models.Task{Id: id}
	err := task.GetTaskByID()
	return task, err
}

func (t taskService) CreateTask(task models.Task) (*models.Task, error) {
	task.Id = uuid.NewMD5(uuid.Nil, []byte(task.Title))
	err := task.Create()
	return &task, err
}

func (t taskService) GetTasks() ([]*models.Task, error) {
	var task models.Task
	return task.GetTasks()
}

func (t taskService) DeleteTask(task models.Task) error {
	return task.Delete()
}

func (t taskService) UpdateTask(task models.Task) (*models.Task, error) {
	err := task.Update()
	return &task, err
}
