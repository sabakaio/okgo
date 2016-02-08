package models

import (
	"encoding/json"
	"github.com/docker/libkv/store"
	"strings"
)

// Job - job definition
type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	When    string `json:"when"`
	Once    bool   `json:"once"`
}

// CreateJob - creates a Job and saves it to KV
func CreateJob(name string, command string, schedule string) (job *Job, err error) {
	once := false
	schedule = strings.TrimSpace(schedule)
	if schedule == "" || strings.Contains(schedule, "once") {
		once = true
	}
	job = &Job{
		Name:    name,
		Command: command,
		Once:    once,
	}

	jsonStr, err := json.Marshal(job)
	if err != nil {
		return
	}
	err = kv.Put("jobs/"+name, jsonStr, &store.WriteOptions{IsDir: true})
	return
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
		json.Unmarshal(job.Value, &_t)
		t = append(t, _t)
	}

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// GetJob - get a job by name
func GetJob(name string) (*Job, error) {
	var t Job
	job, err := kv.Get("jobs/" + name)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(job.Value, &t)

	if err != nil {
		return nil, err
	}

	return &t, nil
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
