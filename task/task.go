package task

import (
	"bufio"
	"fmt"
	"os"
)

func AddTask(tasks *[]string) {
	fmt.Println("--------------------------------")
	fmt.Printf("Enter task name: ")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	*tasks = append(*tasks, taskName)

	fmt.Println("Your updated tasks list:")
	for index, task := range *tasks {
		fmt.Printf("%v - %v\n", index+1, task)
	}
}

func DeleteTask(tasks *[]string) {
	fmt.Println("--------------------------------")
	fmt.Printf("Enter task index: ")
	var taskIndex int
	fmt.Scanln(&taskIndex)
	if taskIndex >= 1 && taskIndex <= len(*tasks) {
		*tasks = append((*tasks)[:taskIndex-1], (*tasks)[taskIndex:]...)

		fmt.Println("These are your update tasks list:")
		for index, task := range *tasks {
			fmt.Printf("%v - %v\n", index+1, task)
		}
	} else {
		fmt.Println("Invalid task index")
	}
}

func ViewTaskList(tasks []string) {
	fmt.Println("--------------------------------")
	if len(tasks) == 0 {
		fmt.Println("You don't have any tasks")
	} else {
		fmt.Println("These are your tasks:")
		for index, task := range tasks {
			fmt.Printf("%v - %v\n", index+1, task)
		}
	}
}
