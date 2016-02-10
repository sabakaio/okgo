package models

import (
	"encoding/json"
	"github.com/docker/libkv/store"
	"strings"
)

type Model interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// Job - job definition
type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	When    string `json:"when"`
	Once    bool   `json:"once"`
}

func NewJob(name string, command string, schedule string) *Job {
	once := false
	schedule = strings.TrimSpace(schedule)
	if schedule == "" || strings.Contains(schedule, "once") {
		once = true
	}
	job := &Job{
		Name:    name,
		Command: command,
		Once:    once,
	}
	return job
}

func UnmarshallJob(data []byte) (*Job, error) {
	job := &Job{}
	err := job.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *Job) Marshal() ([]byte, error) {
	data, err := json.Marshal(*j)
	return data, err
}

func (j *Job) Unmarshal(data []byte) error {
	return json.Unmarshal(data, j)
}

// CreateJob - creates a Job and saves it to KV
func CreateJob(name string, command string, schedule string) (*Job, error) {
	job := NewJob(name, command, schedule)
	data, err := job.Marshal()
	if err != nil {
		return nil, err
	}
	err = kv.Put("jobs/"+name, data, &store.WriteOptions{IsDir: true})
	return job, nil
}

// ListJobs - returns all jobs
func ListJobs() (*[]Job, error) {
	var t []Job
	jobs, err := kv.List("jobs")
	if err != nil {
		return &t, nil
	}

	for _, job := range jobs {
		var _t = Job{}
		_t.Unmarshal(job.Value)
		t = append(t, _t)
	}

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// GetJob - get a job by name
func GetJob(name string) (job *Job, err error) {
	pair, err := kv.Get("jobs/" + name)
	if err != nil {
		return
	}
	return UnmarshallJob(pair.Value)
}

// RemoveJob - removes a job by name
func RemoveJob(name string) error {
	err := kv.Delete("jobs/" + name)
	if err != nil {
		return err
	}
	return nil
}

// PurgeJobs - removes a job by name
func PurgeJobs() error {
	err := kv.DeleteTree("jobs/")
	if err != nil {
		return err
	}
	return nil
}
