package main

import "fmt"

func main() {
	// y := 17
	var (
		x int
		y string
		z bool
	)
	x = 16
	y = "B4LBU"
	z = true

	// fmt.Printf("Mon age est: %v", x)
	// fmt.Printf("\nMon prenom est: %v", y)
	// fmt.Printf("\nJe suis dev: %v", z)
	fmt.Printf("Mon age est : %v.\nMon prenom est : %v.\nJe suis dev: %v", x, y, z)
}
