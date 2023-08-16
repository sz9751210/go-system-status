package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {

	color.Cyan("Welcome to System Status Monitor CLI!")
	color.Cyan("------------------------------------")

	for {
		clearScreen()

		choice, err := chooseStatus()
		if err != nil {
			color.Red("Error: %v", err)
			break
		}

		switch choice {
		case "Memory Status":
			printMemoryStatus()
		case "CPU Status":
			printCPUStatus()
		case "Disk Status":
			printDiskStatus()
		case "Operating System":
			printOSInfo()
		case "Kernel Version":
			printKernelVersion()
		case "Uptime":
			printUptime()
		case "Shell":
			printShell()
		case "OS Distribution":
			printOSDistro()
		case "Quit":
			color.Cyan("Exiting...")
			return
		default:
			color.Red("Invalid choice.")
		}

		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()

		clearScreen()

	}
}
