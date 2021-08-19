package taskHandler

import (
	"net/http"
	"task-list/task"

	"github.com/gin-gonic/gin"
)

var tasks = []task.Task{
	{Title: "Test", Content: "This is the content", Id: "1"},
}

// getTasks responds with the list of all albums as JSON.
func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// postTask adds an album from JSON received in the request body.
func PostTask(c *gin.Context) {
	var newTask task.Task

	// Call BindJSON to bind the received JSON to
	// newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	// Add the new album to the slice.
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range tasks {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
