package main

import (
	"fmt"

	"github.com/Balbul/my-module/functions"
)

func CalculateValue(intChannel chan int) {
	randomNumber := functions.GenerateRandomNumber(50)
	intChannel <- randomNumber
}

func main() {
	foo := make(chan int)
	defer close(foo)

	go CalculateValue(foo)

	fmt.Println(foo)
	fmt.Println(<-foo)
}
