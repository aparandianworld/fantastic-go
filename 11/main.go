package main

import (
	"fmt"
	"os"
)

type magicEntry struct {
	offset      int
	bytes       []byte
	description string
}

var magicDB = []magicEntry{
	{0, []byte{0xFF, 0xD8, 0xFF}, "JPEG image"},
}

func magicFile(path string) (string, error) {

	if path == "" {
		return "", fmt.Errorf("path is empty")
	}

	info, err := os.Stat(path)

	if err != nil {
		return "", fmt.Errorf("failed to get file info: %v", err)
	}

	if info.IsDir() {
		return "", fmt.Errorf("path is a directory")
	}

	fh, err := os.Open(path)

	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	defer fh.Close()

	buf := make([]byte, 512)

	_, err = fh.Read(buf)

	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// TO DO: implement magic file logic

	return "", nil

}

func main() {
	var path string

	fmt.Println("Enter path to file: ")
	fmt.Scan(&path)

	interpretation, err := magicFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file magic: ", interpretation)

}
