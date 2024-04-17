package utils

import (
	"fmt"
	"memo/models"

	"github.com/jedib0t/go-pretty/v6/table"
)


func DrawTable(tasks []models.Task) {
	// Create a writer
	tableWriter := table.NewWriter()
	// Append table header
	tableWriter.AppendHeader(table.Row{"ID", "Task", "Deadline", "Completed"})
	// Loop through results and append data to table
	if len(tasks) < 1 {
		tableWriter.AppendRow(table.Row{"You're all caught up for now."})
		tableWriter.AppendRow(table.Row{"Enjoy your free time, or add new tasks to stay productive!"})
	} else {
		for _, task := range tasks {
			tableWriter.AppendRow(table.Row{task.ID, task.TaskName, task.Deadline.Format("2006-01-02 15:04"), task.Completed})
		}
	}
	// Render table
	fmt.Println(tableWriter.Render())
}
