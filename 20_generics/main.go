package main

import "fmt"

// go v1.18+
// type parameters | inference | set

type Number interface {
	int | ~float64 | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8 | float32
}

func main() {
	type f float64
	var foo f = 1.0

	fmt.Println(min[float32](0.7, 0.5)) // type parameters
	fmt.Println(max(foo, 0.5))          // type inference
	fmt.Println(equals(foo, 0.5))       // type set

}

func min[T int32 | float32](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T ~float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func equals[T Number](x, y T) bool {
	return x == y
}
