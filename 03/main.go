package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Banner...")
	message := "Now is better than never."
	width := 32
	padding := max(0, (width-len(message))/2)

	fmt.Println(strings.Repeat(" ", padding) + strings.Repeat("-", len(message)) + strings.Repeat(" ", padding))
	fmt.Println(strings.Repeat(" ", padding) + message + strings.Repeat(" ", padding))
	fmt.Println(strings.Repeat(" ", padding) + strings.Repeat("-", len(message)) + strings.Repeat(" ", padding))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
