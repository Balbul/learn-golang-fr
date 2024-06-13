package main

import "fmt"

func updateA(number int) {
	number = 5
}

func main() {
	number := 10
	updateA(number) // ne modifie pas la valeur
	fmt.Println(number)
}
