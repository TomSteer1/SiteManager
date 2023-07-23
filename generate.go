package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
)

var navTemplate string 
func generatePortfolio(id string) {
	navTemplate = loadTemplate("nav.html")
	portfolio := loadPortfolio(id)
	clear()
	fmt.Println("Generating portfolio for", portfolio.Name)
	createBase(&portfolio)
	generateIndex(&portfolio)
	generateCategories(&portfolio)
	generateProjects(&portfolio)
	generateEducation(&portfolio)
	generateContact(&portfolio)
	fmt.Println("Done!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func createBase(portfolio *Portfolio) {
	os.Mkdir("output", 0777)
	os.Mkdir("output/"+portfolio.Id, 0777)
  exec.Command("cp", "-r", "site/style", "output/"+portfolio.Id).Run()
	exec.Command("cp", "-r", "site/assets", "output/"+portfolio.Id).Run()
	exec.Command("cp", "-r", "site/robots.txt", "output/"+portfolio.Id).Run()
}

func generateIndex(portfolio *Portfolio) {
	indexTemplate := loadTemplate("index.html")
	indexTemplate = strings.Replace(indexTemplate, "##NAME##", portfolio.Name,-1)
	os.WriteFile("output/"+portfolio.Id+"/index.html", []byte(indexTemplate), 0777)
}


func generateCategories(portfolio *Portfolio) {
	categories := ""
	for _, category := range portfolio.Categories {
		fmt.Println("Generating category", category.Name)
		catTemplate := loadTemplate("category.html")
		skills := ""
		for _, skill := range category.Skills {
			skillTemplate := loadTemplate("skill.html")
			skillTemplate = strings.Replace(skillTemplate, "##NAME##", portfolio.Skills[skill].Name,-1)
			skillTemplate = strings.Replace(skillTemplate, "##IMAGE##",portfolio.Skills[skill].Image,-1)
			skillTemplate = strings.Replace(skillTemplate, "##FG##", portfolio.Skills[skill].Fg,-1)
			skillTemplate = strings.Replace(skillTemplate, "##BG##", portfolio.Skills[skill].Bg,-1)
			skills += skillTemplate + "\n"
		}
		catTemplate = strings.Replace(catTemplate, "##NAME##", category.Name,-1)
		catTemplate = strings.Replace(catTemplate, "##SKILLS##", skills,-1)
		if len(category.Skills) > 0 && category.Id != "hidden" {
			categories += catTemplate + "\n"
		}
	}
	output := loadTemplate("skills.html")
	output = strings.Replace(output, "##CATEGORIES##", categories,-1)
	output = strings.Replace(output, "##NAME##", portfolio.Name,-1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/skills.html", []byte(output), 0777)
}

func generateProjects(portfolio *Portfolio) {
	projects := ""
	for _, project := range portfolio.Projects {
		fmt.Println("Generating project", project.Name)
		projTemplate := loadTemplate("project.html")
		skills := ""
		for _, skill := range project.Skills {
			skillTemplate := loadTemplate("project-tag.html")
			skillTemplate = strings.Replace(skillTemplate, "##NAME##", portfolio.Skills[skill].Name,-1)
			skillTemplate = strings.Replace(skillTemplate, "##FG##", portfolio.Skills[skill].Fg,-1)
			skillTemplate = strings.Replace(skillTemplate, "##BG##", portfolio.Skills[skill].Bg,-1)
			skills += skillTemplate + "\n"
		}
		projTemplate = strings.Replace(projTemplate, "##NAME##", project.Name,-1)
		projTemplate = strings.Replace(projTemplate, "##SKILLS##", skills,-1)
		projTemplate = strings.Replace(projTemplate, "##IMAGE##", project.Image,-1)
		projTemplate = strings.Replace(projTemplate, "##DESCRIPTION##", project.Description,-1)
		projects += projTemplate + "\n"
	}
	output := loadTemplate("projects.html")
	output = strings.Replace(output, "##PROJECTS##", projects,-1)
	output = strings.Replace(output, "##NAME##", portfolio.Name,-1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/projects.html", []byte(output), 0777)
}

func generateEducation(portfolio *Portfolio) {
	output := loadTemplate("education.html")
	output = strings.Replace(output, "##NAME##", portfolio.Name,-1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/education.html", []byte(output), 0777)
}

func generateContact(portfolio *Portfolio) {
	output := loadTemplate("contact.html")
	output = strings.Replace(output, "##NAME##", portfolio.Name,-1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/contact.html", []byte(output), 0777)
}

func loadTemplate(name string) string {
	file, err := os.Open("site/templates/"+name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	var output string
	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}
	return output
}
