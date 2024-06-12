package main

import (
	"errors"
	"fmt"
)

func main() {
	sayHello("B4LBU")
	r := calculatePerimRect(5, 2.3)
	fmt.Printf("r: %v\n", r)

	message, err := sayBye("")

	fmt.Println(message)
	fmt.Println(err)
}

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v.\n", name)
}

func calculatePerimRect(lo, la float64) float64 {
	return 2 * (lo + la)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("\bErreur: Vous avez oublié de spécifier un nom")
	}

	byeMessage := fmt.Sprintf("%v s'en va! Bonne soirée!", name)
	return byeMessage, nil
}
