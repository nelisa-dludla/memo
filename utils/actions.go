package utils

import (
	"bufio"
	"fmt"
	"memo/database"
	"memo/models"
	"os"
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
	var taskName, userInput string
	var taskDate time.Time
	var task models.Task
	color.Cyan("***** Memo - Add Task *****\n\n")
	// Sentence reader
	reader := bufio.NewReader(os.Stdin)
	// Retrieve title
	fmt.Printf("Enter task: ")
	taskName, err := reader.ReadString('\n')
	if err != nil {
		color.Yellow("Invalid input. Input should be a string.")
		return
	}
	// Retrieve deadline
	fmt.Printf("Is this task to be completed today? (Y/N): ")
	_, err = fmt.Scanf("%s", &userInput)
	if err != nil {
		color.Yellow("Invalid input. Input should be a string.")
		return
	}
	// Added information to task model
	task.TaskName = taskName
	task.Completed = false

	if strings.ToLower(userInput) == "y" {
		now := time.Now()
		date := now.Format("2006-01-02")
		// Check if the date format has been entered correctly
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			color.Yellow("Invalid input. Date should be in the format YYYY-MM-DD.")
		}
		taskDate = parsedDate
	} else if strings.ToLower(userInput) == "n" {
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
	parsedTime, err := time.Parse("15-04", timeStr)
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

func EditTask() {
	color.Cyan("***** Memo - Edit Task *****\n\n")
	// Find task
	var task models.Task
	var taskId int
	fmt.Print("Enter task id: ")
	_, err := fmt.Scanf("%d", &taskId)
	if err != nil {
		color.Yellow("Invalid input. Task id should be an integer.")
	}
	result := database.DB.First(&task, taskId)
	if result.Error != nil {
		color.Yellow("Task not found.")
		return
	}

	var userInput int
	fmt.Print(`
(1) Edit Task Name
(2) Edit Deadline

(9) Cancel

> `)

	_, err = fmt.Scanf("%d", &userInput)
	if err != nil {
		color.Yellow("Invalid Input.")
	}

	switch userInput {
	case 1:
		EditTaskName(task)
	case 2:
		EditDeadline(task)
	case 9:
		return
	default:
		color.Yellow("Invalid Input. Enter (1) to Edit Title, (2) to Edit Deadline, or (9) to Cancel.")
	}
}

func EditTaskName(task models.Task) {
	// Sentence reader
	reader := bufio.NewReader(os.Stdin)
	// Retrieve new title
	var newTaskName string
	fmt.Print("New title: ")
	newTaskName, _ = reader.ReadString('\n')

	task.TaskName = newTaskName
	database.DB.Save(&task)

	color.Green("Title changed successfully.")
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
	parsedTime, err := time.Parse("15-04", timeStr)
	fmt.Println("parsedTime:", parsedTime)
	if err != nil {
		color.Yellow("Invalid input. Time should be in the format HH-MM.")
	}

	newDeadline := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
	task.Deadline = newDeadline
	database.DB.Save(&task)

	color.Green("Deadline changed successfully.")
}

func DeleteTask() {
	color.Cyan("***** Memo - Delete Task *****\n\n")
	// Find task
	var task models.Task
	var taskId int
	fmt.Print("Enter task id: ")
	_, err := fmt.Scanf("%d", &taskId)
	if err != nil {
		color.Yellow("Invalid input. Task id should be an integer.")
	}

	result := database.DB.First(&task, taskId)
	if result.Error != nil {
		color.Yellow("Task not found.")
		return
	}
	fmt.Println("The value of result:", result)
	// Delete task from database
	database.DB.Delete(&task)
	// Success message
	color.Green("Task deleted successfully.")
}

func CompletedTask() {
	color.Cyan("***** Memo - Mark Task as Completed *****\n\n")
	// Retrieve task id
	var taskId int
	fmt.Print("Enter task id: ")
	_, err := fmt.Scanf("%d", &taskId)
	if err != nil {
		color.Yellow("Invalid input. Task id should be an integer.")
	}
	// Find task
	var task models.Task
	result := database.DB.Find(&task, taskId)
	if result.Error != nil {
		color.Yellow("Task not found.")
		return
	}
	task.Completed = true
	database.DB.Save(&task)

	color.Green("Task #%d has been completed.", taskId)
}
