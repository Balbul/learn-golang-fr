package main

import "fmt"

func main() {
	var (
		x int
		y int
	)
	x = 15
	y = 15

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	fmt.Println("-------")

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	fmt.Println("-------")
	fmt.Println(x == y && x != y)
	fmt.Println(x == y || x != y)
}
