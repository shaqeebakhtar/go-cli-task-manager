package main

import (
	"example/cli-task-manager/task"
	"fmt"
)

func main() {
	fmt.Println("Welcome to CLI Task Manager Application")

	tasks := []string{}

	for {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Choose an action:")
		fmt.Println("1. Add a new task")
		fmt.Println("2. Delete a task")
		fmt.Println("3. View task list")
		fmt.Println("4. Exit")

		// get user action
		var userAction string
		fmt.Scanln(&userAction)

		switch userAction {
		case "1":
			task.AddTask(&tasks)
		case "2":
			task.DeleteTask(&tasks)
		case "3":
			task.ViewTaskList(tasks)
		case "4":
			fmt.Println("Thanks for using the application")
			return
		default:
			fmt.Println("Invalid action")
		}
	}

}
