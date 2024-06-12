package main

import "fmt"

func main() {
	var (
		b  bool
		s  string
		i  int
		u  uint
		u8 uint8
		i8 int8
		f  float32
	)
	b = true
	s = "B4LBU"
	i = -5
	u = 15   // > 0
	u8 = 254 // 0 - 255
	i8 = 127 // -128 - 127
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, f)
}
