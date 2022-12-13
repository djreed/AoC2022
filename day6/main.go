package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "input"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileContents, err := io.ReadAll(readFile)
	if err != nil {
		panic(err)
	}

	input := string(fileContents)
	fmt.Printf("Input string: %s\n", input)

	// Start of Packet
	REQUIRED_LENGTH := 4
	packetStartIndex := 0
	for i := 0; i < len(input); i++ {
		packetStartIndex = i + 1

		// If we have less than 4 characters in our recent history, wait til we do
		if i < (REQUIRED_LENGTH - 1) {
			continue
		} else {
			startOfBlock := i - (REQUIRED_LENGTH - 1)
			endOfBlock := i + 1

			window := input[startOfBlock:endOfBlock]

			// No duplicate characters in the previous 4 characters
			if containsDuplicates(window) {
				continue
			} else {
				packetStartIndex = i

				fmt.Printf("No duplicates found in character block '%s', at position %d\n", window, packetStartIndex+1)
				break
			}
		}
	}

	REQUIRED_LENGTH = 14
	messageStartIndex := 0
	for i := packetStartIndex; i < len(input); i++ {
		messageStartIndex = i + 1

		// If we have less than 4 characters in our recent history, wait til we do
		if i < (REQUIRED_LENGTH - 1) {
			continue
		} else {
			startOfBlock := i - (REQUIRED_LENGTH - 1)
			endOfBlock := i + 1

			window := input[startOfBlock:endOfBlock]

			// No duplicate characters in the previous 4 characters
			if containsDuplicates(window) {
				continue
			} else {
				messageStartIndex = i

				fmt.Printf("No duplicates found in character block '%s', at position %d\n", window, messageStartIndex+1)
				break
			}
		}
	}
}

func containsDuplicates(str string) bool {
	charMap := make(map[rune]int)
	for _, c := range str {
		if charMap[c] != 0 {
			return true
		} else {
			charMap[c]++
		}
	}
	return false
}
