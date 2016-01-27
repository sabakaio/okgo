package models

import (
	"encoding/json"
	"github.com/docker/libkv/store"
)

// Task - task definition
type Task struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

// CreateTask - creates a Task and saves it to KV
func CreateTask(name string, command string) (*Task, error) {
	task := Task{name, command}
	jsonStr, err := json.Marshal(task)

	if err != nil {
		return nil, err
	}

	kv.Put("tasks/"+name, jsonStr, &store.WriteOptions{IsDir: true})
	return &task, nil
}

// ListTasks - returns all tasks
func ListTasks() (*[]Task, error) {
	var t []Task
	tasks, err := kv.List("tasks")
	if err != nil {
		return &t, nil
	}

	for _, task := range tasks {
		var _t = Task{}
		json.Unmarshal(task.Value, &_t)
		t = append(t, _t)
	}

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// GetTask - get a task by name
func GetTask(name string) (*Task, error) {
	var t Task
	task, err := kv.Get("tasks/" + name)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(task.Value, &t)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// RemoveTask - removes a task by name
func RemoveTask(name string) error {
	err := kv.Delete("tasks/" + name)
	if err != nil {
		return err
	}
	return nil
}

// PurgeTasks - removes a task by name
func PurgeTasks() error {
	err := kv.DeleteTree("tasks/")
	if err != nil {
		return err
	}
	return nil
}
