package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName string
	Deadline time.Time
	Completed bool
}
