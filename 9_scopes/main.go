package main

import "fmt"

var x int

func main() {

	x = 5
	fmt.Println(x)
	f()

	showY() // go run main.go hello.go
}

func f() {
	x := 10
	fmt.Println(x)
}
