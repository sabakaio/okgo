package models

import (
	"encoding/json"
	"github.com/docker/libkv/store"
)

// Token - Token model
type Token struct {
	Name  string `json:"name"`
	Token string `json:null`
}

// CreateJob - creates a Job and saves it to KV
func CreateJob(name string, command string) (*Job, error) {
	job := Job{name, command}
	jsonStr, err := json.Marshal(job)

	if err != nil {
		return nil, err
	}

	kv.Put("jobs/"+name, jsonStr, &store.WriteOptions{IsDir: true})
	return &job, nil
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
