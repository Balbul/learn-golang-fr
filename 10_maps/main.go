package main

import "fmt"

func main() {
	supermarketPrice := map[string]int{
		"prince": 8,
		"eau":    2,
		"meat":   6,
	}

	fmt.Println(supermarketPrice)
	fmt.Println(supermarketPrice["prince"])

	for key, value := range supermarketPrice {
		fmt.Println(key, value)
	}
}
