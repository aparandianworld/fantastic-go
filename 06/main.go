package main

import (
	"fmt"
)

func printCollatzSequence(n int) {
	fmt.Printf("%v -> ", n)
}

func collatz(n int) {

	if n < 0 {
		fmt.Println("Invalid input")
		return
	}

	if n == 0 {
		fmt.Println("0")
		return
	}

	if n == 1 {
		fmt.Println("1")
		return
	}

	if n%2 == 0 {
		printCollatzSequence(n)
		collatz(n / 2)
	} else {
		printCollatzSequence(n)
		collatz(n*3 + 1)
	}
}

func main() {

	fmt.Println("Enter a number: ")
	var num int
	fmt.Scan(&num)

	fmt.Println("Collatz sequence:")
	collatz(num)
}
