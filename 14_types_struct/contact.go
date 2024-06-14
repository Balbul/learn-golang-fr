package main

import "fmt"

type contact struct {
	name    string
	age     int
	numbers map[string]string
}

func newContact(name string, age int, numbers map[string]string) contact {
	return contact{
		name:    name,
		age:     age,
		numbers: numbers,
	}
}

func (c contact) displayContact() string {
	display := fmt.Sprintf("Contact: %v (%v ans)", c.name, c.age)
	display += "\n---------------\n"

	for key, value := range c.numbers {
		display += fmt.Sprintf("%v %v\n", key+":", value)
	}

	return display
}
