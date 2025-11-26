package main

import (
	"fmt"
)

type Addable interface {
	int | float64 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(add(2, 3))
	fmt.Println(add(2.9, 3.4))
	fmt.Println(add("Go ", "rokcs!"))
}
