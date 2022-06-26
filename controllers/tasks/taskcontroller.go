package tasks

import (
	"errors"
	"gin-gonic-tut/httperror"
	"gin-gonic-tut/models"
	"gin-gonic-tut/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if createdTask, err := services.TaskService.CreateTask(task); err == nil {
		c.JSON(http.StatusCreated, createdTask)
	}
}

func UpdateTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httperror.NewBadRequestError("invalid task id"))
	}

	var updateTask *models.Task
	err = c.ShouldBindJSON(&updateTask)
	if err != nil {
		c.JSON(http.StatusNotFound, httperror.NewNotFoundError(err.Error()))
	}
	updateTask.Id = taskID

	_, err = services.TaskService.GetTaskByID(updateTask.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	updatedTask, err := services.TaskService.UpdateTask(*updateTask)
	if err != nil {
		c.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httperror.NewBadRequestError("invalid task id"))
	}

	task, err := services.TaskService.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, httperror.NewNotFoundError(err.Error()))
	}

	err = services.TaskService.DeleteTask(*task)
	if err != nil {
		c.JSON(http.StatusOK, "")
	}
}

func OptionsTask(c *gin.Context) {
	c.Header("allow", "get, post, put, patch, delete, options")
	c.Header("x-token", "asdasdasd-asdasdas-sdfsdf-werwer")
	c.JSON(http.StatusOK, "")
}

func GetTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid task id"))
	}

	if task, err := services.TaskService.GetTaskByID(taskID); err != nil {
		c.JSON(http.StatusNotFound, httperror.NewNotFoundError(err.Error()))
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func GetTasks(c *gin.Context) {
	if task, err := services.TaskService.GetTasks(); err == nil {
		c.JSON(http.StatusOK, task)
	}
}
