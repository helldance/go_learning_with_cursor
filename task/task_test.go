package task

import (
	"testing"
)

func TestNew(t *testing.T) {
	name := "Test Task"
	id := 1
	task := New(name, id)

	if task.Name != name {
		t.Errorf("Expected task name %s, but got %s", name, task.Name)
	}

	if task.ID != id {
		t.Errorf("Expected task ID %d, but got %d", id, task.ID)
	}

	if task.Status != "In Progress" {
		t.Errorf("Expected initial status 'In Progress', but got %s", task.Status)
	}
}

func TestUpdateStatus(t *testing.T) {
	task := New("Test Task", 1)
	newStatus := "Completed"
	task.UpdateStatus(newStatus)

	if task.Status != newStatus {
		t.Errorf("Expected status %s, but got %s", newStatus, task.Status)
	}
}

func TestListTasks(t *testing.T) {
	tasks := make(map[int]Task)
	task1 := New("Task 1", 1)
	task2 := New("Task 2", 2)
	tasks[1] = *task1
	tasks[2] = *task2

	err := ListTasks(tasks)
	if err != nil {
		t.Errorf("Error listing tasks: %v", err)
	}
}
