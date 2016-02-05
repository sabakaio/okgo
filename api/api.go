package api

import (
	"../models"
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
		v1.GET("/jobs", getJobs)
		v1.GET("/jobs/:name", getJob)
		v1.POST("/jobs", createJob)
		v1.DELETE("/jobs/:name", deleteJob)
	}
	return
}

func getJobs(c *gin.Context) {
	jobs, err := models.ListJobs()
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, *jobs)
}

func getJob(c *gin.Context) {
	job, err := models.GetJob(c.Param("name"))
	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, *job)
}

func createJob(c *gin.Context) {
	var t models.Job
	if c.BindJSON(&t) == nil {
		models.CreateJob(t.Name, t.Command)
	} else {
		c.JSON(400, "Bad request")
	}
	c.JSON(200, t)
}

func deleteJob(c *gin.Context) {
	err := models.RemoveJob(c.Param("name"))
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "ok")
}
