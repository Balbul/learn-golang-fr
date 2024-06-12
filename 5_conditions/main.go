package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, "est un nombre pair")
	} else {
		fmt.Println(x, "est un nombre impair")
	}

	y := rand.Int()
	if y%2 == 0 {
		fmt.Println(y, "est un nombre pair")
	} else {
		fmt.Println(y, "est un nombre impair")
	}

	age := 19
	if age > 18 {
		fmt.Println("Je suis majeur")
	} else if age == 18 {
		fmt.Println("Je viens d'être majeur")
	} else {
		fmt.Println("Je suis mineur")
	}
}
