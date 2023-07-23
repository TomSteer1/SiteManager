package main

import (
	"fmt"
	"os"
)

const version = "0.0.2"
const author = "Tom Steer"
var configDir string

func init() {
	clear()
	configDir = os.Getenv("HOME") + "/.config/tomsteer/portfolio-manager"
	// Check if config directory exists
	_, err := os.Stat(configDir)
	if err != nil {
		// Create config directory
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			fmt.Println("Error creating config directory: " + err.Error())
			os.Exit(1)
		}
	}
	// Check if portfolios directory exists
	_, err = os.Stat(configDir + "/portfolios")
	if err != nil {
		// Create portfolios directory
		err = os.Mkdir(configDir + "/portfolios", 0755)
		if err != nil {
			fmt.Println("Error creating portfolios directory: " + err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	for true {
		mainMenu()
	}
}

func mainMenu(error ...string) {
	clear()
	if error != nil {
		fmt.Println("Error: " + error[0])
	}
	fmt.Println("1. Create a new portfolio")
	fmt.Println("2. Edit/View a portfolio")
	fmt.Println("3. Generate portfolio")
	fmt.Println("4. Exit")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		createPortfolio()
	case 2:
		id := choosePortfolio()
		for editPortfolio(id) {}
	case 3:
		generatePortfolio(choosePortfolio())
	case 4:
		exit()
	default:
		mainMenu("Invalid choice")
	}
}



func reloadPosts(portfolio *Portfolio) {
}
