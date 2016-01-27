package api

import (
	"github.com/evindor/okgo/models"
	"github.com/gin-gonic/gin"
)

// CreateServer - create a HTTP Api Server
func CreateServer() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	v1 := r.Group("api/v1")
	{
		v1.GET("/tasks", getTasks)
		v1.GET("/tasks/:name", getTask)
		v1.POST("/tasks", createTask)
		v1.DELETE("/tasks/:name", deleteTask)
	}
	return
}

func getTasks(c *gin.Context) {
	tasks, err := models.ListTasks()
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, *tasks)
}

func getTask(c *gin.Context) {
	task, err := models.GetTask(c.Param("name"))
	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, *task)
}

func createTask(c *gin.Context) {
	var t models.Task
	if c.BindJSON(&t) == nil {
		models.CreateTask(t.Name, t.Command)
	} else {
		c.JSON(400, "Bad request")
	}
	c.JSON(200, t)
}

func deleteTask(c *gin.Context) {
	err := models.RemoveTask(c.Param("name"))
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "ok")
}
