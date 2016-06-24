package registry

import (
	"github.com/sabakaio/okgo/models"
)

type JobsRegistry struct {
	Registry
}

func NewJobsRegistry() *JobsRegistry {
	return &JobsRegistry{
		Registry{
			store:      models.GetKv(),
			namePrefix: "jobs",
		},
	}
}

// Create and put into storage new Job model
func (r *JobsRegistry) Create(name string, command string, schedule string) (*models.Job, error) {
	job := models.NewJob(name, command, schedule)
	err := r.Put(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

// Get a model from registry by name
func (r *JobsRegistry) Get(name string) (job *models.Job, err error) {
	value, err := r.get(name)
	if err != nil {
		return
	}
	return models.UnmarshalJob(value)
}

// Get a list of model from registry
func (r *JobsRegistry) List() ([]*models.Job, error) {
	var jobs []*models.Job
	values, err := r.list()
	if err != nil {
		return jobs, err
	}
	for _, value := range values {
		job, err := models.UnmarshalJob(*value)
		if err != nil {
			return make([]*models.Job, 0), err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
