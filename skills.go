package main 

import (
	"fmt"
	"os"
	"bufio"
)

func editSkills(portfolio *Portfolio) bool {
	clear()
	fmt.Println("Editing skills for portfolio " + portfolio.Name)
	fmt.Println("1. Add Skill")
	fmt.Println("2. Remove Skill")
	fmt.Println("3. Edit Skill")
	fmt.Println("4. Add Category")
	fmt.Println("5. Remove Category")
	fmt.Println("6. Edit Category")
	fmt.Println("7. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		addSkill(portfolio)
	case 2:
		skill := chooseSkill(portfolio)
		if skill != nil {
			removeSkill(portfolio, skill)
		}
	case 3:
		skill := chooseSkill(portfolio)
		if skill != nil {
			for editSkill(skill){
				portfolio.Skills[skill.Id] = *skill
				savePortfolio(portfolio)
			}
		}
	case 4:
		addCategory(portfolio)
	case 5:
		category := chooseCategory(portfolio)
		if category != nil {
			removeCategory(portfolio, category)
		}
	case 6:
		category := chooseCategory(portfolio)
		if category != nil {
			for editCategory(category) {
				savePortfolio(portfolio)
			}
		}
	case 7:
		return false
	}
	return true
}

func chooseSkill(portfolio *Portfolio) *Skill {
	clear()
	fmt.Println("Choose a skill")
	var tempSkills []string
	for _, skill := range portfolio.Skills {
		tempSkills = append(tempSkills, skill.Id)
	}	
	for i, skill := range tempSkills {
		fmt.Printf("%d. %s\n", i+1, portfolio.Skills[skill].Name)
	}
	if len(tempSkills) == 0 {
		fmt.Println("No skills found")
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(tempSkills) || choice < 0 {
		chooseSkill(portfolio)
	}
	if choice == 0 {
		return nil 
	}
	skill := portfolio.Skills[tempSkills[choice-1]]
	return &skill
}

func chooseCategory(portfolio *Portfolio) *Category {
	clear()
	fmt.Println("Choose a category")
	for i, category := range portfolio.Categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	if len(portfolio.Skills) == 0 {
		fmt.Println("No categories found")
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(portfolio.Categories) || choice < 0 {
		chooseCategory(portfolio)
	}
	if choice == 0 {
		return nil 
	}
	return &portfolio.Categories[choice-1]
}

func addSkill(portfolio *Portfolio) {
	clear()
	fmt.Println("Adding skill to portfolio " + portfolio.Name)
	var category *Category
	// Pick Category
	fmt.Println("Choose a category")
	for i, category := range portfolio.Categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	fmt.Println("0. New Category")
	var choice int
	input(&choice)
	if choice > len(portfolio.Categories) || choice < 0 {
		addSkill(portfolio)
	}
	if choice == 0 {
		category = addCategory(portfolio)
	} else {
		category = &portfolio.Categories[choice-1]
	}
	newSkill(portfolio,category)
	savePortfolio(portfolio)
}

func newSkill(portfolio *Portfolio, category *Category) *Skill{
	clear()
	fmt.Println("Adding skill to category " + category.Name)
	var skill Skill
	input(&skill.Id, "Skill ID")
	for skill.Id == "" {
		clear()
		fmt.Println("Skill ID cannot be empty")
		input(&skill.Id, "Skill ID")
	}
	if _, ok := portfolio.Skills[skill.Id]; ok {
		clear()
		fmt.Println("Skill ID already exists")
		bufio.NewReader(os.Stdin).ReadString('\n')
		skill = portfolio.Skills[skill.Id]
		return &skill
	}
	input(&skill.Name, "Skill Name")
	for skill.Name == "" {
		clear()
		fmt.Println("Skill Name cannot be empty")
		input(&skill.Name, "Skill Name")
	}
	input(&skill.Fg, "Foreground Color")
	for len(skill.Fg) != 6 {
		clear()
		fmt.Println("Foreground color must be 6 characters long")
		input(&skill.Fg, "Foreground Color")
	}
	input(&skill.Bg, "Background Color")
	for len(skill.Bg) != 6 {
		clear()
		fmt.Println("Background color must be 6 characters long")
		input(&skill.Bg, "Background Color")
	}
	input(&skill.Image, "Image Path")
	for skill.Image == "" {
		clear()
		fmt.Println("Image path cannot be empty")
		input(&skill.Image, "Image Path")
	}
	category.Skills = append(category.Skills, skill.Id)
	portfolio.Skills[skill.Id] = skill
	fmt.Println("Skill added")
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
	return &skill
}

func removeSkill(portfolio *Portfolio, skill *Skill, skip ...bool) {
	clear()
	fmt.Println("Removing skill " + skill.Name + " from portfolio " + portfolio.Name)
	delete(portfolio.Skills, skill.Id)
	for i, category := range portfolio.Categories {
		for j, skillId := range category.Skills {
			if skillId == skill.Id {
				portfolio.Categories[i].Skills = append(category.Skills[:j], category.Skills[j+1:]...)
			}
		}
	}
	for i, project := range portfolio.Projects {
		for j, skillId := range project.Skills {
			if skillId == skill.Id {
				portfolio.Projects[i].Skills = append(project.Skills[:j], project.Skills[j+1:]...)
			}
		}
	}
	savePortfolio(portfolio)	
	if len(skip) == 0 {
		fmt.Println("Skill removed")
		fmt.Println("Press enter to continue")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

func editSkill(skill *Skill) bool {
	clear()
	fmt.Println("Editing skill " + skill.Name)
	var choice int
	fmt.Println("1. ID - " + skill.Id)
	fmt.Println("2. Name - " + skill.Name)
	fmt.Println("3. Foreground Color - " + skill.Fg)
	fmt.Println("4. Background Color - " + skill.Bg)
	fmt.Println("5. Image Path - " + skill.Image)
	fmt.Println("0. Back")
	input(&choice)
	switch choice {
	case 1:
		clear()
		input(&skill.Id, "Skill ID")
		for skill.Id == "" {
			clear()
			fmt.Println("Skill ID cannot be empty")
			input(&skill.Id, "Skill ID")
		}
	case 2:
		clear()
		input(&skill.Name, "Skill Name")
		for skill.Name == "" {
			clear()
			fmt.Println("Skill Name cannot be empty")
			input(&skill.Name, "Skill Name")
		}
	case 3:
		clear()
		input(&skill.Fg, "Foreground Color")
		for len(skill.Fg) != 6 {
			clear()
			fmt.Println("Foreground color must be 6 characters long")
			input(&skill.Fg, "Foreground Color")
		}
	case 4:
		clear()
		input(&skill.Bg, "Background Color")
		for len(skill.Bg) != 6 {
			clear()
			fmt.Println("Background color must be 6 characters long")
			input(&skill.Bg, "Background Color")
		}
	case 5:
		clear()
		input(&skill.Image, "Image Path")
		for skill.Image == "" {
			clear()
			fmt.Println("Image path cannot be empty")
			input(&skill.Image, "Image Path")
		}
	case 0:
		return false
	}
	return true
}

func addCategory(portfolio *Portfolio) *Category{
	var tempCategory Category
	input(&tempCategory.Id, "Category ID")
	for tempCategory.Id == "" {
		clear()
		fmt.Println("Category ID cannot be empty")
		input(&tempCategory.Id, "Category ID")
	}
	input(&tempCategory.Name, "Category Name")
	for tempCategory.Name == "" {
		clear()
		fmt.Println("Category Name cannot be empty")
		input(&tempCategory.Name, "Category Name")
	}
	portfolio.Categories = append(portfolio.Categories, tempCategory)
	savePortfolio(portfolio)
	clear()
	fmt.Println("Category added")
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
	return &portfolio.Categories[len(portfolio.Categories)-1]
}

func removeCategory(portfolio *Portfolio, category *Category) {
	clear()
	var choice string
	input(&choice, "Are you sure you want to remove category " + category.Name + "? (y/n)")
	if choice == "y" {
		for _, skill := range category.Skills {
			sk := portfolio.Skills[skill]
			removeSkill(portfolio, &sk,true)
		}
		for i, cat := range portfolio.Categories {
			fmt.Println(cat.Id)
			fmt.Println(category.Id)
			if cat.Id == category.Id {
				portfolio.Categories = append(portfolio.Categories[:i], portfolio.Categories[i+1:]...)
				break
			}
		}
		fmt.Println(portfolio.Categories)
		savePortfolio(portfolio)
		fmt.Println("Category removed")
	} else {
		fmt.Println("Category not removed")
	}
		fmt.Println("Press enter to continue")
		bufio.NewReader(os.Stdin).ReadString('\n')
}

func editCategory(category *Category) bool {
	clear()
	fmt.Println("Editing category " + category.Name)
	var choice int
	fmt.Println("1. ID - " + category.Id)
	fmt.Println("2. Name - " + category.Name)
	fmt.Println("0. Back")
	input(&choice)
	switch choice {
		case 1:
			clear()
			input(&category.Id, "Category ID")
			for category.Id == "" {
				clear()
				fmt.Println("Category ID cannot be empty")
				input(&category.Id, "Category ID")
			}
		case 2:
			clear()
			input(&category.Name, "Category Name")
			for category.Name == "" {
				clear()
				fmt.Println("Category Name cannot be empty")
				input(&category.Name, "Category Name")
			}
		case 0:
			return false
		}
	return true
}

