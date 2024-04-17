package main

import (
	"memo/database"
	"memo/utils"
)

func init() {
	database.Database()
}

func main() {
	for {
		utils.MainMenu()
	}
}
