package main

import "fmt"

func main() {
	var list [2]int

	list[0] = 10
	list[1] = 20

	newList := [...]int{40, 50}
	fmt.Println(len(newList))
	bigList := [...]int{40, 50, 60, 430, 789012, 501934}

	for k, v := range bigList {
		fmt.Printf("Psition %d est égal à %d.\n", k, v)
	}

}
