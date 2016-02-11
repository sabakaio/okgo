package models

import (
	"encoding/json"
	"strings"
)

type Model interface {
	GetName() string
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

func UnmarshalJob(data []byte) (*Job, error) {
	job := &Job{}
	err := job.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *Job) GetName() string {
	return j.Name
}

func (j *Job) Marshal() ([]byte, error) {
	data, err := json.Marshal(*j)
	return data, err
}

func (j *Job) Unmarshal(data []byte) error {
	return json.Unmarshal(data, j)
}
