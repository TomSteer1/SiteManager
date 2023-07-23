package main

import (
	"fmt"
	"os"
	"bufio"
)

func chooseProject(portfolio *Portfolio) (*Project) {
	clear()
	fmt.Println("Choose a project")
	for i, project := range portfolio.Projects {
		fmt.Printf("%d. %s\n", i+1, project.Name)
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(portfolio.Projects) || choice < 0 {
		chooseProject(portfolio)
	}
	if choice == 0 {
		editProjects(portfolio)
		return nil
	}
	return &portfolio.Projects[choice-1]
}

func editProjects(portfolio *Portfolio) bool {
	clear()
	fmt.Println("Editing projects")
	fmt.Println("1. Add project")
	fmt.Println("2. Edit project")
	fmt.Println("3. Delete project")
	fmt.Println("4. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		addProject(portfolio)
	case 2:
		project := chooseProject(portfolio)
		if project != nil {
			for editProject(portfolio, project) {}
		}
	case 3:
		project := chooseProject(portfolio)
		if project != nil {
			deleteProject(portfolio, project)
		}
	case 4:
		return false
	}
	return true
}

func addProject(portfolio *Portfolio) {
	clear()
	fmt.Println("Adding project")
	var project Project
	input(&project.Id, "Project ID")
	for project.Id == "" {
		clear()
		fmt.Println("ID cannot be empty")
		input(&project.Id, "Project ID")
	}
	input(&project.Name, "Name")
	for project.Name == "" {
		clear()
		fmt.Println("Name cannot be empty")
		input(&project.Name, "Name")
	}
	input(&project.Description, "Description")
	for project.Description == "" {
		clear()
		fmt.Println("Description cannot be empty")
		input(&project.Description, "Description")
	}
	input(&project.Image, "Image")
	for project.Image == "" {
		clear()
		fmt.Println("Image cannot be empty")
		input(&project.Image, "Image")
	}
	portfolio.Projects = append(portfolio.Projects, project)
	savePortfolio(portfolio)
	fmt.Print("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
	var choice string
	input(&choice,"Add skills? (y/n)") 
	if choice == "y" {
		for editProjectSkills(portfolio,&project) {}
	}
}

func editProjectSkills(portfolio *Portfolio, project *Project) bool {
	clear()

	fmt.Println("Editing skills for project " + project.Name)
	fmt.Println("Current skills:")
	for _, skill := range project.Skills {
		fmt.Println("  " + portfolio.Skills[skill].Name)
	}
	if len(project.Skills) == 0 {
		fmt.Println("  None")
	}
	fmt.Println("1. Add skill")
	fmt.Println("2. Remove skill")
	fmt.Println("3. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		addProjectSkill(portfolio,project)
	case 2:
		deleteProjectSkill(portfolio,project)
	case 3:
		return false
	}
	return true
}

func addProjectSkill(portfolio *Portfolio, project *Project) {
	clear()
	fmt.Println("Adding skill to project " + project.Name)
	fmt.Println("Select category")
	for i, category := range portfolio.Categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	fmt.Println("0. New category")
	var choice int
	input(&choice)
	if choice > len(portfolio.Categories) || choice < 0 {
		addProjectSkill(portfolio,project)
	}
	var category *Category
	if choice == 0 {
		category = addCategory(portfolio)
	}
	category = &portfolio.Categories[choice-1]
	fmt.Println("Select skill")
	var tempSkills []Skill
	for i, skill := range category.Skills {
		fmt.Printf("%d. %s\n", i+1, portfolio.Skills[skill].Name)
		tempSkills = append(tempSkills, portfolio.Skills[skill])
	}
	fmt.Println("0. New skill")
	input(&choice)
	if choice > len(category.Skills) || choice < 0 {
		addProjectSkill(portfolio,project)
	}
	var skill *Skill
	if choice == 0 {
		skill = newSkill(portfolio,category)
	} else{	
		skill = &tempSkills[choice-1]
	}
	project.Skills = append(project.Skills, skill.Id)
	savePortfolio(portfolio)
}

func deleteProjectSkill(portfolio *Portfolio, project *Project) {
}

func editProject(portfolio *Portfolio, project *Project) bool {
	clear()
	fmt.Println("Editing project " + project.Name + " - " + project.Description)
	fmt.Println("1. ID - " + project.Id)
	fmt.Println("2. Name - " + project.Name)
	fmt.Println("3. Description - " + project.Description)
	fmt.Println("4. Image - " + project.Image)
	var skillsString string
	for _, skill := range project.Skills {
		skillsString += portfolio.Skills[skill].Name + ", "
	}
	if len(project.Skills) > 0 {
		skillsString = skillsString[:len(skillsString)-2]
	} else{
		skillsString = "None"
	}
	fmt.Println("5. Skills - " + skillsString)
	fmt.Println("6. Back")
	var choice int
	input(&choice)
	clear()
	switch choice {
	case 1:
		input(&project.Id, "Project ID")
		for project.Id == "" {
			clear()
			fmt.Println("ID cannot be empty")
			input(&project.Id, "Project ID")
		}
	case 2:
		input(&project.Name, "Name")
		for project.Name == "" {
			clear()
			fmt.Println("Name cannot be empty")
			input(&project.Name, "Name")
		}
	case 3:
		input(&project.Description, "Description")
		for project.Description == "" {
			clear()
			fmt.Println("Description cannot be empty")
			input(&project.Description, "Description")
		}
	case 4:
		input(&project.Image, "Image")
		for project.Image == "" {
			clear()
			fmt.Println("Image cannot be empty")
			input(&project.Image, "Image")
		}
	case 5:
		for editProjectSkills(portfolio,project) {}
	case 6:
		return false
	}
	savePortfolio(portfolio)
	return true
}

func deleteProject(portfolio *Portfolio, project *Project) {
}
