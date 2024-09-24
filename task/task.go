package task

import (
	"encoding/json"
	"fmt"
)

// Task represents a task with a name, status, and ID
// add json tags to the struct
type Task struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	ID     int    `json:"id"`
}

// New creates a new Task with the given name and ID, and default status "In Progress"
func New(name string, id int) *Task {
	return &Task{
		Name:   name,
		Status: "In Progress",
		ID:     id,
	}
}

// new function to create a task from json
func NewFromJSON(jsonString string) (*Task, error) {
	var task Task
	err := json.Unmarshal([]byte(jsonString), &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateStatus updates the task's status
func (t *Task) UpdateStatus(newStatus string) {
	t.Status = newStatus
}

// customize print task to print line by line
func (t *Task) String() string {
	return fmt.Sprintf("ID: %d\nName: %s\nStatus: %s\n", t.ID, t.Name, t.Status)
}

// ListTasks prints all tasks in JSON format as a single object
func ListTasks(tasks map[int]Task) error {
	jsonTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks to JSON: %w", err)
	}
	fmt.Println("All tasks in JSON format:")
	fmt.Println(string(jsonTasks))
	return nil
}
