package main

import (
	"fmt"
	"hello/task"
)

func main() {
	// change it to a map
	tasks := make(map[int]task.Task)
	// Test task
	task1 := task.New("Complete Go project", 1)
	task2 := task.New("Complete Web project", 2)
	task3, err := task.NewFromJSON(`{"name": "Complete Mobile project", "status": "In Progress", "id": 3}`)
	if err != nil {
		fmt.Println("Error creating task from JSON:", err)
		return
	}

	tasks[1] = *task1
	tasks[2] = *task2
	tasks[3] = *task3

	err2 := task.ListTasks(tasks)
	if err2 != nil {
		fmt.Println("Error listing tasks:", err2)
		return
	}

	// update task status
	task1.UpdateStatus("Completed")
	fmt.Printf("Updated Task: %+v\n", task1)
}
