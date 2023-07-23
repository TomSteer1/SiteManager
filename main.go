package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

const version = "0.0.1"
const author = "Tom Steer"
var configDir string

type Portfolio struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Projects []Project `json:"projects"`
	Categories []Category `json:"categories"`
	Skills map[string]Skill `json:"skills"`
	Education []Education `json:"education"`
	Contacts []Contact `json:"contacts"`
	Pages Pages `json:"pages"`
}

type Project struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Skills []string `json:"skills"`
	Image string `json:"image"`
}

type Category struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Skills []string `json:"skills"`
}

type Skill struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Fg string `json:"fg"`
	Bg string `json:"bg"`
}

type Education struct {
	Id string
	Name string
	Degree string
	Start string
	End string
	Years []Year
}

type Year struct {
	Id string
	Name string
	Modules []Module
}

type Module struct {
	Id string
	Name string
	Grade string
	Link string
	Description string
}

type Contact struct {
	Id string
	Name string
	Username string
	Link string
	Icon string
}

type Pages struct {
	Home bool `json:"home"`
	Skills bool `json:"skills"`
	Projects bool `json:"projects"`
	Education bool `json:"education"`
	Posts bool `json:"posts"`
	Contact bool `json:"contact"`
}

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
	fmt.Println("7. Reload Posts")
	fmt.Println("8. Back")
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
		editEducation(&portfolio)
	case 6:
		editContacts(&portfolio)
	case 7:
		reloadPosts(&portfolio)
	case 8:
		return false
	case -1:
		deletePortfolio(&portfolio)
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


func editEducation(portfolio *Portfolio) {
}

func editContacts(portfolio *Portfolio) {
}

func reloadPosts(portfolio *Portfolio) {
}

func deletePortfolio(portfolio *Portfolio) {
}



// Helper Functions


func exit() {
	clear()
	fmt.Println("Goodbye")
	os.Exit(0)
}

func clear() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Portfolio Manager v" + version + " by " + author)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}

func input(variable interface{},prompt ...string) {
	if prompt != nil {
		fmt.Println(prompt[0])
	}
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	handleError(err)
	switch variable.(type) {
	case *string:
		*variable.(*string) = line[:len(line)-1]
	case *int:
		if len(line) == 1 {
			*variable.(*int) = 0
		} else {
			*variable.(*int), err = strconv.Atoi(line[:len(line)-1])
			if(err != nil) {
				*variable.(*int) = 0
			}
		}
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
