package main

import (
	"github.com/fatih/color"
	"learnGO/menu"
)

func main() {
	color.Green("__Password manager__")
App:
	for {
		breaked := menu.ChooseMenu()
		if breaked {
			break App
		}
	}
}
