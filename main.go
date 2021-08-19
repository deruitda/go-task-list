package main

import (
	"task-list/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskById)
	router.POST("/tasks", handlers.PostTask)

	router.Run("localhost:8080")
}
