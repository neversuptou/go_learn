package main

import (
	"learnGO/menu"
)

func main() {
App:
	for {
		breaked := menu.ChooseMenu()
		if breaked {
			break App
		}
	}
}
