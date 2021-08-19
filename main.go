package main

import (
	"task-list/taskHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", taskHandler.GetTasks)
	router.GET("/tasks/:id", taskHandler.GetTaskById)
	router.POST("/tasks", taskHandler.PostTask)

	router.Run("localhost:8080")
}
