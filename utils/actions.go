package utils

import (
	"fmt"
	"memo/database"
	"memo/models"
	"strings"
	"time"
)

func RetrieveAllTasks() []models.Task {
	var tasks []models.Task

	result := database.DB.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Application Error: Could not retrieve tasks from database.")
	}
	return tasks
}

func AddTask() {
	var title, description, deadline_opt1 string
	var taskDate time.Time
	var task models.Task
	// Retrieve title
	fmt.Printf("Enter a task title: ")
	_, err := fmt.Scanf("%s", &title)
	if err != nil {
		fmt.Println("Invalid input. Input should be a string.")
	}
	// Retrieve description
	fmt.Printf("Enter a task description: ")
	_, err = fmt.Scanf("%s", &description)
	if err != nil {
		fmt.Println("Invalid input. Input should be a string.")
	}
	// Retrieve deadline
	fmt.Printf("Is this task to be completed today? (Y/N): ")
	_, err = fmt.Scanf("%s", &deadline_opt1)
	if err != nil {
		fmt.Println("Invalid input. Input should be a string.")
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
			fmt.Println("Invalid input. Date should be in the format YYYY-MM-DD.")
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
			fmt.Println("Invalid input. Date should be in the format YYYY-MM-DD.")
		}
		taskDate = parsedDate
	} else {
		fmt.Println("Invalid input.")
	}
	// Retrieve time
	var timeStr string
	fmt.Printf("Enter a time - (Format - HH-MM): ")
	_, err = fmt.Scanf("%s", &timeStr)
	if err != nil {
		fmt.Println("Invalid input. Time should be in the format HH-MM.")
	}
	// Check if the time format has been enter correctly
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		fmt.Println("Invalid input. Time should be in the format HH-MM.")
	}

	deadline := time.Date(taskDate.Year(), taskDate.Month(), taskDate.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), parsedTime.Nanosecond(), parsedTime.Location())
	task.Deadline = deadline

	// Add task to database
	result := database.DB.Create(&task)
	if result.Error != nil {
		fmt.Println("Application Error. Failed to add task.")
	}
}
