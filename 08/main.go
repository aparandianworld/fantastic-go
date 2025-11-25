package main

import (
	"fmt"
)

func sum(values []float64) float64 {
	total := 0.0

	for _, val := range values {
		total += val
	}

	return total
}

func mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("empty values")
	}

	return (sum(values) / float64(len(values))), nil
}

func main() {
	fmt.Println("Mean value of [1.2, 2.4, 3.6, 4.8, 5.0] is: ")

	values := []float64{1.2, 2.4, 3.6, 4.8, 5.0}
	mv, err := mean(values)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(mv)
}
