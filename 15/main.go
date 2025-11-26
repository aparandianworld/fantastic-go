package main

import (
	"fmt"
)

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(add(2, 3))
	fmt.Println(add(2.9, 3.4))
	fmt.Println(add("Go ", "rokcs!"))
}
