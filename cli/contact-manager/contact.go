package main

import (
	"errors"
	"strings"
)

type Contact struct {
	Name  string
	Email string
	Phone string
}

func (c Contact) Summary() string {
	return c.Name + " | " + c.Email + " | " + c.Phone
}

func findContact(contacts []Contact, name string) (Contact, error) {
	for _, contact := range contacts {
		if strings.EqualFold(contact.Name, name) {
			return contact, nil
		}
	}
	return Contact{}, errors.New("Contact not found!")
}
