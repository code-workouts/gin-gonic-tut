package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"gin-gonic-tut/cache"
	"gin-gonic-tut/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"type:time" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"type:time" json:"updated_at,omitempty"`
}

var (
	db = repositories.DB
)

func init() {
	err := db.AutoMigrate(&Task{})
	if err != nil {
		panic(err)
	}
	//populateTestData()
}

// Populates initial temp data to database
func populateTestData(){
	vmwareHcmpTaskUUID := uuid.NewMD5(uuid.Nil, []byte("VMware HCMP"))
	log.Println("vmwareHcmpTaskUUID: ", vmwareHcmpTaskUUID)

	vmwareMigrationTaskUUID := uuid.NewMD5(uuid.Nil, []byte("VMware Migration"))
	log.Println("vmwareMigrationTaskUUID: ", vmwareMigrationTaskUUID)

	vmwareApprovalTaskUUID := uuid.NewMD5(uuid.Nil, []byte("VMware Approval"))
	log.Println("vmwareApprovalTaskUUID: ", vmwareMigrationTaskUUID)

	tasks := []Task{
		{
			Id:          vmwareHcmpTaskUUID,
			Title:       "VMware HCMP",
			Description: "VMware HCMP Service",
		},
		{
			Id:          vmwareMigrationTaskUUID,
			Title:       "VMware Migration",
			Description: "VMware Migration Service",
		},
		{
			Id:          vmwareApprovalTaskUUID,
			Title:       "VMware Approval",
			Description: "VMware Approval Service",
		},
	}
	db.Create(&tasks)
}

func (task *Task) getTaskByIDFromCache() (bool, error) {
	if err := cache.Ping(); err != nil {
		return false, nil
	}

	value, err := cache.Get(task.Id.String())
	if err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		return false, err
	}
	if value != nil {
		b := bytes.NewReader(value)
		if err := json.NewDecoder(b).Decode(task); err != nil {
			return false, err
		}
		log.Println("found entry in cache")
		return true, err
	}
	return false, nil
}

func (task *Task) setTaskByIDToCache() error {
	b := bytes.Buffer{}
	if err := json.NewEncoder(&b).Encode(task); err != nil {
		return err
	}
	return cache.Set(task.Id.String(), b.Bytes(), cache.DefaultExpiration)
}

func (task *Task) replaceTaskByIDOnCache() error {
	b := bytes.Buffer{}
	if err := json.NewEncoder(&b).Encode(task); err != nil {
		return err
	}
	return cache.Replace(task.Id.String(), b.Bytes(), cache.DefaultExpiration)
}

func (task *Task) deleteTaskByIDOnCache() error {
	b := bytes.Buffer{}
	if err := json.NewEncoder(&b).Encode(task); err != nil {
		return err
	}
	 return cache.Delete(task.Id.String())
}

func (task *Task) GetTaskByID() error {
	found, err := task.getTaskByIDFromCache()
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	log.Println("entry in cache not found")
	// Note: This is to mimic delay
	time.Sleep(time.Second * 1 / 2)
	result := db.First(task, "id = ?", task.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return task.setTaskByIDToCache()
}

func (task *Task) Create() error {
	result := db.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return task.setTaskByIDToCache()
}

func (task *Task) CreateInBatches(tasks []*Task) error {
	result := db.CreateInBatches(&tasks, 5)
	if result.Error != nil {
		return result.Error
	}

	for _, nTask := range tasks {
		err := nTask.setTaskByIDToCache()
		if err != nil {
			return err
		}
	}
	return nil
}

func (task *Task) Update() error {
	result := db.Save(&task)
	if result.Error != nil {
		return result.Error
	}

	return task.replaceTaskByIDOnCache()
}

func (task *Task) Delete() error {
	result := db.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	return task.deleteTaskByIDOnCache()
}

func (task *Task) findByTitle() error {
	result := db.First(&task, "title = ?", task.Title)
	if result.Error != nil {
		return result.Error
	}
	return task.setTaskByIDToCache()
}
