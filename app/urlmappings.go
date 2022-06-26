package app

import (
	"gin-gonic-tut/controllers/ping"
	"gin-gonic-tut/controllers/redirect"
	"gin-gonic-tut/controllers/tasks"
)

func mapurls() {
	v1 := router.Group("/v1")

	// Ping controller routes
	v1.GET("/ping", ping.Ping)

	// Tasks Controller routes
	v1.GET("/tasks/:id", tasks.GetTask)
	v1.GET("/tasks", tasks.GetTasks)
	v1.POST("/tasks", tasks.CreateTask)
	v1.PUT("/tasks/:id", tasks.UpdateTask)
	v1.DELETE("/tasks/:id", tasks.DeleteTask)
	v1.OPTIONS("/tasks", tasks.OptionsTask)

	// Test Redirection
	v1.GET("/google", redirect.Redirect)
}
