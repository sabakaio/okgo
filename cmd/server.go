package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/evindor/okgo/models"
)

// CmdServer - start a server
var CmdServer = cli.Command{
	Name:        "server",
	Usage:       "start okgo server",
	Description: `start okgo server`,
	Action:      serverAction,
}

func serverAction(c *cli.Context) {
	r := gin.Default()
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
	r.Run(":3000")
	println("Started server on port 3000")
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
