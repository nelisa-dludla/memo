package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string
	Description string
	Deadline time.Time
	Completed bool
}
