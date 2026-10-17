package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		return input
	}
}

func addContact(reader *bufio.Reader, contacts []Contact) []Contact {
	fmt.Println("\n--- Add Contact ---")

	name := readInput(reader, "Name: ")
	email := readInput(reader, "Email: ")
	phone := readInput(reader, "Phone: ")

	contact := Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	contacts = append(contacts, contact)

	fmt.Println("Contact added successfully!")

	return contacts

}

func listContacts(contacts []Contact) {
	fmt.Println("\n--- Contact List ---")

	if len(contacts) == 0 {
		fmt.Println("No contacts found")
		return
	}

	for i, contact := range contacts {
		fmt.Printf("%d. %s\n", i+1, contact.Summary())
	}
}

func searchContact(reader *bufio.Reader, contacts []Contact) {
	fmt.Println("\n--- Find Contact ---")

	name := readInput(reader, "Enter name: ")

	contact, err := findContact(contacts, name)

	if err != nil {
		fmt.Printf("Errpr:", err)
		return
	}

	fmt.Println("Contact found:")
	fmt.Println(contact.Summary())
}

func printMenu() {
	fmt.Println("\n=== Contact Manager ===")
	fmt.Println("1. Add Contact")
	fmt.Println("2. List Contacts")
	fmt.Println("3. Find Contact")
	fmt.Println("4. Exit")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	contacts := []Contact{}

	for {
		printMenu()

		choice := readInput(reader, "Choose an option: ")

		switch choice {
		case "1":
			contacts = addContact(reader, contacts)

		case "2":
			listContacts(contacts)

		case "3":
			searchContact(reader, contacts)

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1-4.")
		}
	}
}
