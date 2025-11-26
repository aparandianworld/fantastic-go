package main

import (
	"fmt"
	"os"
)

func readFirstBytes(filename string, size int) ([]byte, error) {

	if filename == "" {
		return nil, fmt.Errorf("filename is empty")
	}

	if size < 0 {
		return nil, fmt.Errorf("size is negative")
	}

	_, err := os.Stat(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	fh, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	defer fh.Close()

	buf := make([]byte, size)
	n, err := fh.Read(buf)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	if n != size {
		return nil, fmt.Errorf("error reading number of bytes: %v", n)
	}

	return buf[:n], nil
}

func main() {

	filename := "./dog.jpeg"
	size := 10

	data, err := readFirstBytes(filename, size)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", data)
}
