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
	devDir := os.Getenv("HOME") + "/.config/tomsteer"
	configDir = os.Getenv("HOME") + "/.config/tomsteer/portfolio-manager"
	ensureDir(devDir)
	ensureDir(configDir)
	ensureDir(configDir + "/portfolios")
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
