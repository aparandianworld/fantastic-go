package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("What is the most common word in the folliwng sentence: the quick brown fox jumps over the lazy dog")

	sentence := "the quick brown fox jumps over the lazy dog"

	sentence = strings.ToLower(sentence)

	words := strings.Split(sentence, " ")
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	fmt.Println("DEBUG: ", wordsMap)

	max := 0
	mostCommon := ""
	mostCommonCount := 0

	for word, count := range wordsMap {
		if count > max {
			max = count
			mostCommon = word
			mostCommonCount = count
		}
	}

	fmt.Println("The most common word is: `", mostCommon, "` with a count of: ", mostCommonCount)

}
