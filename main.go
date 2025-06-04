package main

import (
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"learnGO/menu"
	"learnGO/utils"
)

func main() {
	color.Green("__Password manager__")
	err := godotenv.Load()
	if err != nil {
		utils.PrintError("Error loading .env file")
	}
App:
	for {
		breaked := menu.ChooseMenu()
		if breaked {
			break App
		}
	}
}
