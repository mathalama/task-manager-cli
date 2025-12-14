package entity

import "fmt"

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks = []Task{}
var nextTaskID = 1

func AddTask(description string) int {
	task := Task{
		ID:          nextTaskID,
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, task)
	return task.ID
}

func UpdateTask(id int, description string, completed bool) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			if description != "" {
				tasks[i].Description = description
			}
			tasks[i].Completed = completed
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func ListAllTasks() {

	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}

	for _, task := range tasks {
		status := "Not completed"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | %s | %s\n", task.ID, task.Description, status)
	}
}
