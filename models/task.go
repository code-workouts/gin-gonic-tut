package models

import (
	"gin-gonic-tut/repositories"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string	     `json:"description"`
	CreatedAt   time.Time    `gorm:"type:time" json:"created_at,omitempty"`
	UpdatedAt   time.Time    `gorm:"type:time" json:"updated_at,omitempty"`
}

var(
	db = repositories.DB
)

func init() {
	db.AutoMigrate(&Task{})
	tasks := []Task{
		{
			Id: uuid.New(),
			Title: "VMware HCMP",
			Description: "VMware HCMP Service",
		},
		{
			Id: uuid.New(),
			Title: "VMware Migration",
			Description: "VMware Migration Service",
		},
		{
			Id: uuid.New(),
			Title: "VMware Approval",
			Description: "VMware Approval Service",
		},
	}
	db.Create(&tasks)
}

func (task *Task) Get() {
	db.First(&task, "title = ?", task.Id)
}

func (task *Task) Create() {
	db.Create(&task)
}

func (task *Task) CreateInBatches(tasks []*Task) {
	db.CreateInBatches(&tasks, 5)
}

func (task *Task) Update() {
}

func (task *Task) Delete() {
}

func (task *Task) findByTitle() {
	db.First(&task, "title = ?", task.Title)
}

func (task *Task) GetTasks() ([]*Task, error) {
	var tasks []*Task
	result := db.Find(&tasks)
	return tasks, result.Error
}
