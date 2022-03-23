package app

import (
	"gin-gonic-tut/controllers/ping"
	"gin-gonic-tut/controllers/redirect"
	"gin-gonic-tut/controllers/tasks"
)

func mapurls() {
	// Ping controller routes
	router.GET("/ping", ping.Ping)

	// Tasks Controller routes
	router.GET("/tasks/:id", tasks.GetTask)
	router.GET("/tasks", tasks.GetTasks)
	router.POST("/tasks", tasks.CreateTask)
	router.PUT("/tasks/:id", tasks.UpdateTask)
	router.DELETE("/tasks/:id", tasks.DeleteTask)

	// Test Redirection
	router.GET("/google", redirect.Redirect)
}
