package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
var contacts []Contact
// Ctrl+Shift+i
func main() {
	loadContacts()
	for {
		dispalyMenu()
		var choice int
		fmt.Print("Select the option:")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input.Please enter the Number")
			continue
		}
		switch choice {
		case 1:
			addContact()
		case 2:
			viewContact()
		case 3:
			searchContacts()
		case 4:
			deleteContact()
		case 5:
			saveContacts()
			fmt.Println("Existing...")
			return
		default:
			fmt.Println("Invalid choice.please select again")
		}
	}
}
func dispalyMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1. Add Contact")
	fmt.Println("2. View Contact")
	fmt.Println("3. Search Contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Exit")
}
func addContact() {
	var name, phone, email string
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Phone Number: ")
	fmt.Scan(&phone)
	fmt.Print("Enter Email: ")
	fmt.Scan(&email)

	contact := Contact{Name: name, PhoneNumber: phone, Email: email}
	contacts = append(contacts, contact)
	saveContacts()
	fmt.Println("Contact added successfully")
}
func viewContact() {
	if len(contacts) == 0 {
		fmt.Println("No contact to display")
		return
	}
	fmt.Println("\nContacts:")
	for i, contact := range contacts {
		fmt.Printf("%d. Name: %s,Phone: %s,Email:%s\n", i+1, contact.Name, contact.PhoneNumber, contact.Email)
	}
}
func searchContacts() {
	var searchTerm string
	fmt.Print("Enter the name or phone number to search:")
	fmt.Scan(&searchTerm)

	fmt.Println("/nSearch Results:")
	found := false
	for _, contact := range contacts {
		if strings.Contains(contact.Name, searchTerm) || strings.Contains(contact.PhoneNumber, searchTerm) {
			fmt.Printf("Name: %s,Phone: %s ,Email:%s\n", contact.Name, contact.PhoneNumber, contact.Email)
			found = true
		}
	}
	if !found {
		fmt.Println("No Contact Found")
	}
}
func deleteContact() {
	var nameOrPhone string
	fmt.Print("Enter the name or phoneNumber of the contact delete")
	fmt.Scan(&nameOrPhone)
	for i, contact := range contacts {
		if contact.Name == nameOrPhone || contact.PhoneNumber == nameOrPhone {
			fmt.Printf("Are you want to delete the contact:%s? (y/n):", contact.Email)
			var confirmation string
			fmt.Scan(&confirmation)
			if strings.ToLower(confirmation) == "y" {
				contacts = append(contacts[:i], contacts[i+1:]...)
				saveContacts()
				fmt.Println("contact deleted successfully")
			} else {
				fmt.Println("Deletion canceled")
			}
			return
		}
	}
	fmt.Println("Contact not Found")
}
func saveContacts() {
	data, err := json.MarshalIndent(contacts, "", " ")
	if err != nil {
		fmt.Println("Error saving contacts:", err)
	}
	err = os.WriteFile("contact.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file ", err)
	}
}
func loadContacts() {
	file, err := os.Open("contact.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error for reading contact file", err)
		return
	}
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		fmt.Println("Error parsing contact JSON:", err)
	}
}
