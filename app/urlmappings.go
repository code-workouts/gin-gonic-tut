package app

import (
	"gin-gonic-tut/controllers/ping"
	"gin-gonic-tut/controllers/tasks"
)

func mapurls() {
	// Ping controller routes
	router.GET("/ping", ping.Ping)

	router.GET("/tasks", tasks.GetTasks)
	router.GET("/tasks/:id", tasks.GetTask)
	router.POST("/tasks", tasks.CreateTask)
}