package main

import (
	"fmt"
	"os"
	"bufio"
)

func editEducation(portfolio *Portfolio) bool {
	clear()
	fmt.Println("Editing education for portfolio " + portfolio.Name)
	fmt.Println("1. Add Education")
	fmt.Println("2. Remove Education")
	fmt.Println("3. Edit Education")
	fmt.Println("4. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		addEducation(portfolio)
	case 2:
		education := chooseEducation(portfolio)
		if education != nil {
			removeEducation(portfolio, education)
		}
	case 3:
		education := chooseEducation(portfolio)
		if education != nil {
			for editEducationStage(education) {
				savePortfolio(portfolio)
			}
		}
	case 4:
		return false
	}
	return true
}

func chooseEducation(portfolio *Portfolio) *Education {
	clear()
	fmt.Println("Choose an education")
	for i, education := range portfolio.Education {
		fmt.Printf("%d. %s\n", i+1, education.Name)
	}
	if len(portfolio.Education) == 0 {
		fmt.Println("No education found")
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(portfolio.Education) || choice < 0 {
		return chooseEducation(portfolio)
	}
	if choice == 0 {
		return nil
	}
	return &portfolio.Education[choice-1]
}	

func addEducation(portfolio *Portfolio) {
	clear()
	fmt.Println("Adding education")
	var education Education
	input(&education.Id, "Education ID")
	input(&education.Name, "Education Name")
	input(&education.File, "Education File")
	for true 	{
	// Check if education file exists
		if _, err := os.Stat("site/templates/education/" + education.File + ".html"); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			input(&education.File, "Education File")
		} else {
			break
		}
	}
	portfolio.Education = append(portfolio.Education, education)
	savePortfolio(portfolio)
}

func removeEducation(portfolio *Portfolio, education *Education) {
	clear()
	fmt.Println("Removing education " + education.Name)
	for i, education := range portfolio.Education {
		if education.Id == education.Id {
			portfolio.Education = append(portfolio.Education[:i], portfolio.Education[i+1:]...)
			savePortfolio(portfolio)
			break
		}
	}
	fmt.Println("Education removed")
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func editEducationStage(education *Education) bool {
	clear()
	fmt.Println("Editing education " + education.Name)
	fmt.Println("1. ID - " + education.Id)
	fmt.Println("2. Name - " + education.Name)
	fmt.Println("3. File - " + education.File)
	fmt.Println("4. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		input(&education.Id, "Education ID")
	case 2:
		input(&education.Name, "Education Name")
	case 3:
		input(&education.File, "Education File")
		for true 	{
			// Check if education file exists
			if _, err := os.Stat("site/templates/education/" + education.File + ".html"); os.IsNotExist(err) {
				fmt.Println("File does not exist")
				input(&education.File, "Education File")
			} else {
				break
			}
		}
	case 4:
		return false
	}
	return true
}
