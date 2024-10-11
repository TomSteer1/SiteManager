package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
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
	fmt.Println("Starting server on port 8080")
	fmt.Println("Press enter to stop and return to menu")
	s := &http.Server{
		Addr:    ":8080",
		Handler: addHtml(http.FileServer(http.Dir("output/" + portfolio.Id))),
	}
	go s.ListenAndServe()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	s.Close()
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
	indexTemplate = strings.Replace(indexTemplate, "##NAME##", portfolio.Name, -1)
	indexTemplate = strings.Replace(indexTemplate, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/index.html", []byte(indexTemplate), 0644)
}

func generateCategories(portfolio *Portfolio) {
	categories := ""
	for _, category := range portfolio.Categories {
		fmt.Println("Generating category", category.Name)
		catTemplate := loadTemplate("category.html")
		skills := ""
		for _, skill := range category.Skills {
			skillTemplate := loadTemplate("skill.html")
			skillTemplate = strings.Replace(skillTemplate, "##NAME##", portfolio.Skills[skill].Name, -1)
			skillTemplate = strings.Replace(skillTemplate, "##IMAGE##", portfolio.Skills[skill].Image, -1)
			skillTemplate = strings.Replace(skillTemplate, "##FG##", portfolio.Skills[skill].Fg, -1)
			skillTemplate = strings.Replace(skillTemplate, "##BG##", portfolio.Skills[skill].Bg, -1)
			skills += skillTemplate + "\n"
		}
		catTemplate = strings.Replace(catTemplate, "##NAME##", category.Name, -1)
		catTemplate = strings.Replace(catTemplate, "##SKILLS##", skills, -1)
		if len(category.Skills) > 0 && category.Id != "hidden" {
			categories += catTemplate + "\n"
		}
	}
	output := loadTemplate("skills.html")
	output = strings.Replace(output, "##CATEGORIES##", categories, -1)
	output = strings.Replace(output, "##NAME##", portfolio.Name, -1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/skills.html", []byte(output), 0644)
}

func generateProjects(portfolio *Portfolio) {
	projects := ""
	for _, project := range portfolio.Projects {
		fmt.Println("Generating project", project.Name)
		projTemplate := loadTemplate("project.html")
		skills := ""
		for _, skill := range project.Skills {
			skillTemplate := loadTemplate("project-tag.html")
			skillTemplate = strings.Replace(skillTemplate, "##NAME##", portfolio.Skills[skill].Name, -1)
			skillTemplate = strings.Replace(skillTemplate, "##FG##", portfolio.Skills[skill].Fg, -1)
			skillTemplate = strings.Replace(skillTemplate, "##BG##", portfolio.Skills[skill].Bg, -1)
			skills += skillTemplate + "\n"
		}
		projTemplate = strings.Replace(projTemplate, "##NAME##", project.Name, -1)
		projTemplate = strings.Replace(projTemplate, "##SKILLS##", skills, -1)
		projTemplate = strings.Replace(projTemplate, "##IMAGE##", project.Image, -1)
		projTemplate = strings.Replace(projTemplate, "##DESCRIPTION##", project.Description, -1)
		projTemplate = strings.Replace(projTemplate, "##LINK##", project.Link, -1)
		projects += projTemplate + "\n"
	}
	output := loadTemplate("projects.html")
	output = strings.Replace(output, "##PROJECTS##", projects, -1)
	output = strings.Replace(output, "##NAME##", portfolio.Name, -1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	os.WriteFile("output/"+portfolio.Id+"/projects.html", []byte(output), 0644)
}

func generateEducation(portfolio *Portfolio) {
	education := ""
	for _, edu := range portfolio.Education {
		fmt.Println("Generating education", edu.Name)
		eduTemplate := loadTemplate("education/" + edu.File + ".html")
		education += eduTemplate + "\n"
	}
	output := loadTemplate("education.html")
	output = strings.Replace(output, "##NAME##", portfolio.Name, -1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	output = strings.Replace(output, "##EDUCATION##", education, -1)
	os.WriteFile("output/"+portfolio.Id+"/education.html", []byte(output), 0644)
}

func generateContact(portfolio *Portfolio) {
	contacts := ""
	for _, contact := range portfolio.Contacts {
		fmt.Println("Generating contact", contact.Name)
		contactTemplate := loadTemplate("contact.html")
		contactTemplate = strings.Replace(contactTemplate, "##NAME##", contact.Name, -1)
		contactTemplate = strings.Replace(contactTemplate, "##IMAGE##", contact.Image, -1)
		contactTemplate = strings.Replace(contactTemplate, "##LINK##", contact.Link, -1)
		contactTemplate = strings.Replace(contactTemplate, "##USERNAME##", contact.Username, -1)
		contacts += contactTemplate + "\n"
	}
	output := loadTemplate("contacts.html")
	output = strings.Replace(output, "##NAME##", portfolio.Name, -1)
	output = strings.Replace(output, "##NAV##", navTemplate, -1)
	output = strings.Replace(output, "##CONTACTS##", contacts, -1)
	os.WriteFile("output/"+portfolio.Id+"/contact.html", []byte(output), 0644)
}

// func generateExperience(portfolio *Portfolio) {
// 	experience := ""
// 	for _, exp := range portfolio.Experience {
// 		fmt.Println("Generating experience", exp.Name)
// 		expTemplate := loadTemplate("experience/" + exp.File + ".html")
// 		experience += expTemplate + "\n"
// 	}
// 	output := loadTemplate("experience.html")
// 	output = strings.Replace(output, "##NAME##", portfolio.Name,-1)
// 	output = strings.Replace(output, "##NAV##", navTemplate, -1)
// 	output = strings.Replace(output, "##EXPERIENCE##", experience,-1)
// 	os.WriteFile("output/"+portfolio.Id+"/experience.html", []byte(output), 0644)
// }

func loadTemplate(name string) string {
	file, err := os.Open("site/templates/" + name)
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

func addHtml(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[len(r.URL.Path)-1] != '/' && !strings.Contains(r.URL.Path, ".") {
			r.URL.Path += ".html"
		}
		next.ServeHTTP(w, r)
	})
}
