package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func MainMenu() {
	color.Cyan("\n***** Memo - A lightweight todo list application *****\n\n")
	//Retrieve tasks from database
	tasks := RetrieveAllTasks()
	DrawTable(tasks)
	// User options
	var userInput int
	fmt.Print(`
(1) Add Task
(2) Edit Task
(3) Delete Task
(4) Mark Task as Completed

(9) Exit Memo

> `)

	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		fmt.Println("Invalid input. Enter (1) to add a task, (2) to edit a task, (3) to delete a task or (9) to Quit Memo.")
	}

	switch userInput {
	case 1:
		AddTask()
	case 2:
		EditTask()
	case 3:
		DeleteTask()
	case 4:
		CompletedTask()
	case 9:
		color.Green("Closing Memo...")
		os.Exit(0)
	default:
		fmt.Println("Invalid input. Enter (1) to add a task, (2) to edit a task, (3) to delete a task or (9) to Quit Memo.")
	}
}
