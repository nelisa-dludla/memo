package utils

import (
	"fmt"
	"memo/database"
	"memo/models"
	"strings"
	"time"

	"github.com/fatih/color"
)

func RetrieveAllTasks() []models.Task {
	var tasks []models.Task

	result := database.DB.Find(&tasks)
	if result.Error != nil {
		color.Red("Application Error: Could not retrieve tasks from database.")
	}
	return tasks
}

func AddTask() {
	var title, description, deadline_opt1 string
	var taskDate time.Time
	var task models.Task
	color.Cyan("***** Memo - Add Task *****\n\n")
	// Retrieve title
	fmt.Printf("Enter a task title: ")
	_, err := fmt.Scanln("%s", &title)
	if err != nil {
		color.Yellow("Invalid input. Input should be a string.")
		return
	}
	// Retrieve description
	fmt.Printf("Enter a task description: ")
	_, err = fmt.Scanln("%s", &description)
	if err != nil {
		color.Yellow("Invalid input. Input should be a string.")
		return
	}
	// Retrieve deadline
	fmt.Printf("Is this task to be completed today? (Y/N): ")
	fmt.Scanln()
	_, err = fmt.Scanf("%s", &deadline_opt1)
	if err != nil {
		color.Yellow("Invalid input. Input should be a string.")
		return
	}
	// Added information to task model
	task.Title = title
	task.Description = description
	task.Completed = false

	if strings.ToLower(deadline_opt1) == "y" {
		now := time.Now()
		date := now.Format("2006-01-02")
		// Check if the date format has been entered correctly
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			color.Yellow("Invalid input. Date should be in the format YYYY-MM-DD.")
		}
		taskDate = parsedDate
	} else if strings.ToLower(deadline_opt1) == "n" {
		// Retrieve date
		var date string
		fmt.Printf("Enter a date - (Format - YYYY-MM-DD): ")
		fmt.Scanf("%s", &date)
		// Check if the date format has been entered correctly
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			color.Yellow("Invalid input. Date should be in the format YYYY-MM-DD.")
		}
		taskDate = parsedDate
	} else {
		color.Yellow("Invalid input.")
	}
	// Retrieve time
	var timeStr string
	fmt.Printf("Enter a time - (Format - HH-MM): ")
	_, err = fmt.Scanf("%s", &timeStr)
	if err != nil {
		color.Yellow("Invalid input. Time should be in the format HH-MM.")
		return
	}
	// Check if the time format has been enter correctly
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		color.Yellow("Invalid input. Time should be in the format HH-MM.")
		return
	}

	deadline := time.Date(taskDate.Year(), taskDate.Month(), taskDate.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
	task.Deadline = deadline

	// Add task to database
	result := database.DB.Create(&task)
	if result.Error != nil {
		color.Red("Application Error. Failed to add task.")
		return
	}

	color.Green("Task was added successfully.")
}

func EditTask(id int) {
	color.Cyan("***** Memo - Edit Task *****\n\n")
	// Find task
	var task models.Task
	result := database.DB.Find(&task, id)
	if result != nil {
		color.Yellow("Task not found.")
		return
	}

	var userInput int
	fmt.Print(`
(1) Edit Title
(2) Edit Description
(3) Edit Deadline

(9) Cancel

> `)

	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		color.Yellow("Invalid Input.")
	}

	switch userInput {
	case 1:
		EditTitle(task)
	case 2:
		EditDescription(task)
	case 3:
		EditDeadline(task)
	case 9:
		return
	default:
		color.Yellow("Invalid Input. Enter (1) to Edit Title, (2) to Edit Description, (3) to Edit Deadline or (9) to Cancel.")
	}
}

func EditTitle(task models.Task) {
	// Retrieve new title
	var newTitle string
	fmt.Print("New title: ")
	fmt.Scanln("%s", &newTitle)

	task.Title = newTitle
	database.DB.Save(&task)

	color.Green("Title changed successfully.")
}

func EditDescription(task models.Task) {
	// Retrieve new description
	var newDescription string
	fmt.Print("New description: ")
	fmt.Scanln("%s", &newDescription)

	task.Description = newDescription
	database.DB.Save(&task)

	color.Green("Description changed successfully.")
}

func EditDeadline(task models.Task) {
	// Retrieve Date
	var date string
	fmt.Printf("Enter a date - (Format - YYYY-MM-DD): ")
	fmt.Scanf("%s", &date)
	// Check if the date format has been entered correctly
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		color.Yellow("Invalid input. Date should be in the format YYYY-MM-DD.")
		return
	}
	// Retrieve time
	var timeStr string
	fmt.Printf("Enter a time - (Format - HH-MM): ")
	_, err = fmt.Scanf("%s", &timeStr)
	if err != nil {
		color.Yellow("Invalid input. Time should be in the format HH-MM.")
		return
	}
	// Check if the time format has been enter correctly
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		color.Yellow("Invalid input. Time should be in the format HH-MM.")
	}

	newDeadline := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
	task.Deadline = newDeadline
	database.DB.Save(&task)

	color.Green("Deadline changed successfully.")
}
