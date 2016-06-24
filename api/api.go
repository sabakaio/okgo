package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sabakaio/okgo/models"
	"github.com/sabakaio/okgo/registry"
)

var jobs *registry.JobsRegistry

func init() {
	jobs = registry.NewJobsRegistry()
}

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
	jobs, err := jobs.List()
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, jobs)
}

func getJob(c *gin.Context) {
	job, err := jobs.Get(c.Param("name"))
	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, *job)
}

func createJob(c *gin.Context) {
	var t models.Job
	if c.BindJSON(&t) == nil {
		jobs.Create(t.Name, t.Command, "")
	} else {
		c.JSON(400, "Bad request")
	}
	c.JSON(200, t)
}

func deleteJob(c *gin.Context) {
	err := jobs.Delete(c.Param("name"))
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "ok")
}
