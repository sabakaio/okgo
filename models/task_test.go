package models

import (
	"testing"
)

func TestCreateJob(t *testing.T) {
	job, err := CreateJob("test", "command")
	if err != nil {
		t.Error(err)
	}
	if job.Name != "test" {
		t.Errorf("error creating job")
	}
}

func TestListJobs(t *testing.T) {
	jobs, err := ListJobs()
	if err != nil {
		t.Error(err)
	}
	if len(*jobs) != 1 {
		t.Errorf("error listing jobs")
	}
}

func TestGetJob(t *testing.T) {
	job, err := GetJob("test")
	if err != nil {
		t.Error(err)
	}
	if job.Command != "command" {
		t.Errorf("error getting job")
	}
}

func TestRemoveJob(t *testing.T) {
	err := RemoveJob("test")
	if err != nil {
		t.Error(err)
	}
	_, err = kv.Get("test")
	if err == nil {
		t.Error("job was not removed")
	}
}
