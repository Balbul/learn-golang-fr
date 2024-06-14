package main

import "fmt"

type Animal interface {
	Noise() string
	NumberOgLegs() int
}

type Cat struct {
	Name  string
	Breed string
}

type Spider struct {
	Name     string
	Breed    string
	Venomous bool
}

func main() {
	cat := Cat{Name: "Kit", Breed: "Main Coon"}
	PrintAnimalInfo(&cat)

	spider := Spider{Name: "Spidy", Breed: "Black widow"}
	PrintAnimalInfo(&spider)
}

func PrintAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.Noise(), "et il possède", a.NumberOgLegs(), "pattes!")
}

func (c *Cat) Noise() string {
	return "Miaow"
}

func (c *Cat) NumberOgLegs() int {
	return 4
}

func (c *Spider) Noise() string {
	return "Hiss"
}

func (c *Spider) NumberOgLegs() int {
	return 8
}
