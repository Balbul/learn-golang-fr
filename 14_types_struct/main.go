package main

import "fmt"

func main() {
	myContact := newContact("B4LBU", 40, map[string]string{"Fixe": "10.11.12.13.14"})
	fmt.Println(myContact)
	fmt.Println(myContact.displayContact())
}
