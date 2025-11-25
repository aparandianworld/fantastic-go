package main

import (
	"fmt"
	"strings"
)

func splitExtension(path string) (string, string) {

	i := strings.LastIndex(path, ".")
	if i == -1 {
		return path, ""
	}
	return path[:i], path[i:]
}

func main() {
	fmt.Println("Enter a file path: ")

	var fp string
	fmt.Scan(&fp)

	if fp == "" {
		fmt.Println("File path is empty")
		return
	}

	root, ext := splitExtension(fp)

	fmt.Println("Root:", root)
	fmt.Println("Extension:", ext)
}
