package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	user := User{
		ID:       1,
		Name:     "Aikyn",
		Email:    "test@test",
		Password: "test",
	}

	tasks := []Task{}
	nextTaskID := 1

	fmt.Println("You need to log in to use the system.")
	var scanEmail string
	fmt.Print("Enter email: ")
	fmt.Scanln(&scanEmail)

	if scanEmail != user.Email {
		fmt.Println("Please try again later")
		return
	}

	fmt.Printf("Welcome to Task Manager, %s!\n", user.Name)

	for {
		fmt.Println("\nChoose what you want to do:")
		fmt.Println("1. Add task")
		fmt.Println("2. Update task")
		fmt.Println("3. Delete task")
		fmt.Println("4. View all tasks")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter: ")
		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {

		case 1:
			var description string
			fmt.Print("Enter your task: ")
			fmt.Scanln(&description)

			task := Task{
				ID:          nextTaskID,
				Description: description,
				Completed:   false,
			}
			nextTaskID++

			tasks = append(tasks, task)
			fmt.Println("Task successfully added")

		case 2:
			var id int
			fmt.Print("Enter task ID: ")
			fmt.Scan(&id)
			fmt.Scanln()

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					found = true

					var updateDescription string
					fmt.Print("New description (press ENTER to skip): ")
					fmt.Scanln(&updateDescription)
					if updateDescription != "" {
						tasks[i].Description = updateDescription
					}

					var status int
					fmt.Print("Completed? (1 = yes, 2 = no): ")
					fmt.Scan(&status)
					fmt.Scanln()

					if status == 1 {
						tasks[i].Completed = true
					} else if status == 2 {
						tasks[i].Completed = false
					}

					fmt.Println("Task updated")
					break
				}
			}

			if !found {
				fmt.Println("Task not found")
			}

		case 3:
			var id int
			fmt.Print("Enter task ID: ")
			fmt.Scan(&id)
			fmt.Scanln()

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					found = true
					fmt.Println("Task deleted")
					break
				}
			}

			if !found {
				fmt.Println("Task not found")
			}

		case 4:
			if len(tasks) == 0 {
				fmt.Println("No tasks")
				continue
			}

			for _, task := range tasks {
				status := "Not completed"
				if task.Completed {
					status = "Completed"
				}
				fmt.Printf("ID: %d | %s | %s\n", task.ID, task.Description, status)
			}

		case 0:
			fmt.Println("Thank you for using Task Manager. See you soon!")
			return

		default:
			fmt.Println("Enter a correct value!")
		}
	}
}
