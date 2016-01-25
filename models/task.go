package models

import (
	"encoding/json"
  "github.com/docker/libkv/store"
)

// Task - task definition
type Task struct {
	Name    string
	Command string
}

// CreateTask - creates a Task and saves it to KV
func CreateTask(name string, command string) (*Task, error) {
	task := Task{name, command}
	jsonStr, err := json.Marshal(task)

	if err != nil {
		return nil, err
	}

	kv.Put("tasks/" + name, jsonStr, &store.WriteOptions{IsDir: true})
	return &task, nil
}

// ListTasks - returns all tasks
func ListTasks() (*[]Task, error) {
	tasks, err := kv.List("tasks")
	if err != nil {
		return nil, err
	}

	var t []Task
	for _, task := range tasks {
		var _t = Task{}
		json.Unmarshal(task.Value, &_t)
		t = append(t, _t)
	}

	if err != nil {
		return nil, err
	}

	return &t, nil;
}
