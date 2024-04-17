package database

import (
	"log"
	"memo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database() {
	// Create/Connect to sqlite database
	db, err := gorm.Open(sqlite.Open("memo.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}
	// Create/Sync tables
	db.AutoMigrate(&models.Task{})
	// Make database variable accessible through the codebase
	DB = db
}
