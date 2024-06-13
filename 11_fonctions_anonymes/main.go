package main

func main() {
	w := func() {
		println("je suis une fonction anonyme sans return")
	}

	z := func() string {
		println("je suis une fonction anonyme avec return")
		return "B4LBU"
	}

	w()
	println(z())
}
