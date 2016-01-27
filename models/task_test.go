package models

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	task, err := CreateTask("test", "command")
	if err != nil {
		t.Error(err)
	}
	if task.Name != "test" {
		t.Errorf("error creating task")
	}
}

func TestListTasks(t *testing.T) {
	tasks, err := ListTasks()
	if err != nil {
		t.Error(err)
	}
	if len(*tasks) != 1 {
		t.Errorf("error listing tasks")
	}
}

func TestGetTask(t *testing.T) {
	task, err := GetTask("test")
	if err != nil {
		t.Error(err)
	}
	if task.Command != "command" {
		t.Errorf("error getting task")
	}
}

func TestRemoveTask(t *testing.T) {
	err := RemoveTask("test")
	if err != nil {
		t.Error(err)
	}
	_, err = kv.Get("test")
	if err == nil {
		t.Error("task was not removed")
	}
}
