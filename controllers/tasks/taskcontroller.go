package tasks

import (
	"gin-gonic-tut/models"
	"gin-gonic-tut/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)
	if createdTask, err := services.TaskService.CreateTask(task); err == nil {
		c.JSON(http.StatusCreated, createdTask)
	}
}

func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	if task, err := services.TaskService.GetTask(taskID); err != nil {
		c.JSON(http.StatusOK, task)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetTasks(c *gin.Context) {
	if task, err := services.TaskService.GetTasks(); err == nil {
		c.JSON(http.StatusOK, task)
	}
}
