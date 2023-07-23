package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"io/ioutil"
)

func createPortfolio() {
	clear()
	var portfolio Portfolio
	fmt.Println("Pick an ID for your portfolio")
	var err error
	for err == nil {
		input(&portfolio.Id)
		_, err = os.Stat(configDir + "/portfolios/" + portfolio.Id)
		if err == nil {
			clear()
			fmt.Println("ID already exists")
			fmt.Println("Pick an ID for your portfolio")
		}
	}
	clear()
	input(&portfolio.Name, "Pick a name for your portfolio")
	for portfolio.Name == "" {
		clear()
		fmt.Println("Name cannot be empty")
		input(&portfolio.Name, "Pick a name for your portfolio")
	}
	clear()
	input(&portfolio.Title, "Pick a title for your portfolio")
	for portfolio.Title == "" {
		clear()
		fmt.Println("Title cannot be empty")
		input(&portfolio.Title, "Pick a title for your portfolio")
	}
	portfolio.Categories = append(portfolio.Categories, Category{Id: "hidden", Name: "Hidden from site (use as tags)"})
	portfolio.Skills = make(map[string]Skill)
	fmt.Println("Creating portfolio...")
	err = os.Mkdir(configDir + "/portfolios/" + portfolio.Id, 0755)
	handleError(err)
	fmt.Println("Creating portfolio config file...")
	file, err := os.Create(configDir + "/portfolios/" + portfolio.Id + "/config")
	handleError(err)
	defer file.Close()
	json, err := json.Marshal(portfolio)
	handleError(err)
	_, err = file.Write(json)
	handleError(err)
	fmt.Println("Portfolio created")
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
	for editPortfolio(portfolio.Id){
	}
}


func choosePortfolio() (string){
	clear()
	fmt.Println("Choose a portfolio")
	files, err := os.ReadDir(configDir + "/portfolios")
	handleError(err)
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}
	if len(files) == 0 {
		fmt.Println("No portfolios found")
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(files) || choice < 0 {
		choosePortfolio()
	}
	if choice == 0 {
		mainMenu()
		exit()
	}
	return files[choice-1].Name()
}


func editPortfolio(id string, err ...string) bool {
	portfolio := loadPortfolio(id)
	clear()
	if err != nil {
		fmt.Println("Error: " + err[0])
	}
	fmt.Println("Editing portfolio " + portfolio.Name)
	fmt.Println("1. Print Portfolio Details")
	fmt.Println("2. Change Title")
	fmt.Println("3. Edit Projects")
	fmt.Println("4. Edit Skills")
	fmt.Println("5. Edit Education")
	fmt.Println("6. Edit Contacts")
//	fmt.Println("7. Reload Posts")
	fmt.Println("7. Back")
	fmt.Println("-1. Delete Portfolio")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		printPortfolio(&portfolio)
	case 2:
		changeTitle(&portfolio)
	case 3:
		for editProjects(&portfolio) {
		}
	case 4:
		for editSkills(&portfolio) {}
	case 5:
		for editEducation(&portfolio) {}
	case 6:
		for editContacts(&portfolio) {}
//	case 7:
//		reloadPosts(&portfolio)
	case 7:
		return false
	case -1:
		deletePortfolio(&portfolio)
		return false
	default:
		return editPortfolio(id, "Invalid choice")
	}
	return true
}

func printPortfolio(portfolio *Portfolio) {
	clear()
	fmt.Println("ID: " + portfolio.Id)
	fmt.Println("Name: " + portfolio.Name)
	fmt.Println("Title: " + portfolio.Title)
	fmt.Println("Projects:")
	for _, project := range portfolio.Projects {
		fmt.Println("  " + project.Name)
	}
	if len(portfolio.Projects) == 0 {
		fmt.Println("  None")
	}
	fmt.Println("Skills:")
	for _, category := range portfolio.Categories {
		fmt.Println("  " + category.Name + ":")
		for _, skill := range category.Skills {
			fmt.Println("    " + portfolio.Skills[skill].Name)
		}
		if len(category.Skills) == 0 {
			fmt.Println("    None")
		}
	}
	fmt.Println("Education:")
	for _, education := range portfolio.Education {
		fmt.Println("  " + education.Name)
	}
	if len(portfolio.Education) == 0 {
		fmt.Println("  None")
	}
	fmt.Println("Contacts:")
	for _, contact := range portfolio.Contacts {
		fmt.Println("  " + contact.Name)
	}
	if len(portfolio.Contacts) == 0 {
		fmt.Println("  None")
	}
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
}



func changeTitle(portfolio *Portfolio) {
	clear()
	fmt.Println("Current title: " + portfolio.Title)
	input(&portfolio.Title, "New title")
	for portfolio.Title == "" {
		clear()
		fmt.Println("Title cannot be empty")
		input(&portfolio.Title, "New title")
	}
	savePortfolio(portfolio)
}



func deletePortfolio(portfolio *Portfolio) {
	clear()
	fmt.Println("Are you sure you want to delete this portfolio? (y/n)")
	var choice string
	input(&choice)
	if choice == "y" {
		err := os.RemoveAll(configDir + "/portfolios/" + portfolio.Id)
		handleError(err)
		fmt.Println("Portfolio deleted")
		fmt.Println("Press enter to continue")
		bufio.NewReader(os.Stdin).ReadString('\n')
	} else {
		editPortfolio(portfolio.Id)
	}
}


func loadPortfolio(id string) Portfolio {
	var portfolio Portfolio
	file, err := os.Open(configDir + "/portfolios/" + id + "/config")
	handleError(err)
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	handleError(err)
	json.Unmarshal([]byte(content), &portfolio)
	fmt.Println(portfolio)
	return portfolio
}

func savePortfolio(portfolio *Portfolio) {
	content, err := json.Marshal(&portfolio)
	handleError(err)
	err = os.WriteFile(configDir + "/portfolios/" + portfolio.Id + "/config",content, 0644)
	handleError(err)
}
