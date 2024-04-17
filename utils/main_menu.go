package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func MainMenu() {
	color.Cyan("***** Memo - A lightweight todo list application *****\n\n")
	//Retrieve tasks from database
	tasks := RetrieveAllTasks()
	DrawTable(tasks)
	// User options
	var userInput int
	fmt.Print(`
(1) Add Task
(2) Edit Task
(3) Delete Task

(9) Exit Memo

> `)

	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		fmt.Println("Invalid input. Enter (1) to add a task, (2) to edit a task, (3) to delete a task or (9) to Quit Memo.")
	}

	switch userInput {
	case 1:
		fmt.Println("Add task")
	case 2:
		fmt.Println("Edit task")
	case 3:
		fmt.Println("Delete task")
	case 9:
		color.Green("Closing Memo...")
		os.Exit(0)
	default:
		fmt.Println("Invalid input. Enter (1) to add a task, (2) to edit a task, (3) to delete a task or (9) to Quit Memo.")
	}
}
