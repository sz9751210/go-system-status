package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func chooseStatus() (string, error) {
	var choice string
	prompt := &survey.Select{
		Message: "Select which status you want to see:",
		Options: []string{
			"Memory Status",
			"CPU Status",
			"Disk Status",
			"Operating System",
			"Kernel Version",
			"Uptime",
			"Shell",
			"OS Distribution",
			"Quit",
		},
	}
	err := survey.AskOne(prompt, &choice)
	return choice, err
}
