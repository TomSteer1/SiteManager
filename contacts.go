package main

import (
	"fmt"
	"os"
	"bufio"
)

func editContacts(portfolio *Portfolio) bool {
	clear()
	fmt.Println("Editing contacts for portfolio " + portfolio.Name)
	fmt.Println("1. Add Contact")
	fmt.Println("2. Remove Contact")
	fmt.Println("3. Edit Contact")
	fmt.Println("4. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		addContact(portfolio)
	case 2:
		contact := chooseContact(portfolio)
		if contact != nil {
			removeContact(portfolio, contact)
		}
	case 3:
		contact := chooseContact(portfolio)
		if contact != nil {
			for editContact(contact) {
				savePortfolio(portfolio)
			}
		}
	case 4:
		return false
	}
	return true	
}

func chooseContact(portfolio *Portfolio) *Contact {
	clear()
	fmt.Println("Choose a contact")
	for i, contact := range portfolio.Contacts {
		fmt.Printf("%d. %s\n", i+1, contact.Name)
	}
	if len(portfolio.Contacts) == 0 {
		fmt.Println("No contacts found")
	}
	fmt.Println("0. Back")
	var choice int
	input(&choice)
	if choice > len(portfolio.Contacts) || choice < 0 {
		return chooseContact(portfolio)
	}
	if choice == 0 {
		return nil
	}
	return &portfolio.Contacts[choice-1]
}

func addContact(portfolio *Portfolio) {
	clear()
	fmt.Println("Adding a new contact")
	var contact Contact
	input(&contact.Id, "ID")
	for contact.Id == "" {
		fmt.Println("ID cannot be empty")
		input(&contact.Id, "ID")
	}
	input(&contact.Name, "Name")
	for contact.Name == "" {
		fmt.Println("Name cannot be empty")
		input(&contact.Name, "Name")
	}
	input(&contact.Username, "Username")
	for contact.Username == "" {
		fmt.Println("Username cannot be empty")
		input(&contact.Username, "Username")
	}
	input(&contact.Link, "Link")
	for contact.Link == "" {
		fmt.Println("Link cannot be empty")
		input(&contact.Link, "Link")
	}
	input(&contact.Image, "Image")
	for contact.Image == "" {
		fmt.Println("Image cannot be empty")
		input(&contact.Image, "Image")
	}
	portfolio.Contacts = append(portfolio.Contacts, contact)
	savePortfolio(portfolio)
}

func removeContact(portfolio *Portfolio, contact *Contact) {
	clear()
	fmt.Println("Removing contact " + contact.Name + " from portfolio " + portfolio.Name)
	for i, c := range portfolio.Contacts {
		if c.Id == contact.Id {
			portfolio.Contacts = append(portfolio.Contacts[:i], portfolio.Contacts[i+1:]...)
			savePortfolio(portfolio)
		}
	}
	fmt.Println("Contact removed")
	fmt.Println("Press enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func editContact(contact *Contact) bool{
	fmt.Println("Editing contact " + contact.Name)
	fmt.Println("1. ID - " + contact.Id)
	fmt.Println("2. Name - " + contact.Name)
	fmt.Println("3. Username - " + contact.Username)
	fmt.Println("4. Link - " + contact.Link)
	fmt.Println("5. Image - " + contact.Image)
	fmt.Println("6. Back")
	var choice int
	input(&choice)
	switch choice {
	case 1:
		input(&contact.Id, "ID")
		for contact.Id == "" {
			fmt.Println("ID cannot be empty")
			input(&contact.Id, "ID")
		}
	case 2:
		input(&contact.Name, "Name")
		for contact.Name == "" {
			fmt.Println("Name cannot be empty")
			input(&contact.Name, "Name")
		}
	case 3:
		input(&contact.Username, "Username")
		for contact.Username == "" {
			fmt.Println("Username cannot be empty")
			input(&contact.Username, "Username")
		}
	case 4:
		input(&contact.Link, "Link")
		for contact.Link == "" {
			fmt.Println("Link cannot be empty")
			input(&contact.Link, "Link")
		}
	case 5:
		input(&contact.Image, "Image")
		for contact.Image == "" {
			fmt.Println("Image cannot be empty")
			input(&contact.Image, "Image")
		}
	case 6:
		return false
	}
	return true
}

