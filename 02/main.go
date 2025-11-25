package main

import (
	"fmt"
)

func main() {
	fmt.Println("Average of all numbers between 1 to 100 that are by 3 or 5...")
	counter, total := 0, 0

	for n := 1 ; n < 101 ; n++ {
		if n % 3 == 0 || n % 5 == 0 {
			counter++
			total += n
		}
	}

	fmt.Println("Average is: ", float64(total)/float64(counter))
}