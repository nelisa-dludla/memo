package database

import (
	"log"
	"memo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	// Create/Connect to sqlite database
	db, err := gorm.Open(sqlite.Open("memo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}
	// Create/Sync tables
	db.AutoMigrate(&models.Task{})
	// Make database variable accessible through the codebase
	DB = db
}
