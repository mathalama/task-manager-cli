package main

import (
	"fmt"
	"task-tracker/entity"
)

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
			id := entity.AddTask(description)
			fmt.Println("Task successfully added. ID:", id)

		case 2:
			var id int
			fmt.Print("Enter task ID: ")
			fmt.Scan(&id)
			fmt.Scanln()

			var desc string
			fmt.Print("New description: ")
			fmt.Scanln(&desc)

			var status int
			fmt.Print("Completed? (1=yes, 2=no): ")
			fmt.Scan(&status)
			fmt.Scanln()

			if !entity.UpdateTask(id, desc, status == 1) {
				fmt.Println("Task not found")
			}

		case 3:
			var id int
			fmt.Print("Enter task ID: ")
			fmt.Scan(&id)
			fmt.Scanln()

			entity.DeleteTask(id)

		case 4:
			entity.ListAllTasks()

		case 0:
			fmt.Println("Thank you for using Task Manager. See you soon!")
			return

		default:
			fmt.Println("Enter a correct value!")
		}
	}
}
