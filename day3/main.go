package main

import (
	"bufio"
	"fmt"
	"os"
)

type Backpack struct {
	c1, c2 string
}

func main() {
	filename := "input"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var backpacks []Backpack
	for fileScanner.Scan() {
		line := fileScanner.Text()
		c1, c2 := splitBackpackString(line)
		backpacks = append(backpacks, Backpack{c1, c2})
	}

	fmt.Printf("Counted %d backpacks\n", len(backpacks))

	var sharedItems []string
	for _, backpack := range backpacks {
		sharedItem := findSharedItem(backpack.c1, backpack.c2)
		sharedItems = append(sharedItems, sharedItem)
	}

	fmt.Printf("Counted %d shared items among backpacks\n", len(sharedItems))

	var itemScores []int
	for _, item := range sharedItems {
		score := getItemScore(item)
		itemScores = append(itemScores, score)
	}

	fmt.Printf("Sum total of shared items: %d\n", Sum(itemScores))

}

func splitBackpackString(contents string) (c1, c2 string) {
	halfLen := len(contents) / 2
	return contents[:halfLen], contents[halfLen:]
}

func findSharedItem(c1, c2 string) string {
	for _, c1Item := range c1 {
		for _, c2Item := range c2 {
			if c1Item == c2Item {
				return string(c1Item)
			}
		}
	}
	return ""
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func getItemScore(item string) int {
	intValue := item[0] // a-z = 97-122, A-Z = 65-90
	if intValue >= 97 && intValue <= 122 {
		intValue -= 96
	} else if intValue >= 65 && intValue <= 90 {
		intValue -= 38
	} else {
		panic("Enexpected back item: " + item)
	}
	return int(intValue)
}

func Sum(list []int) int {
	result := 0
	for _, v := range list {
		result += v
	}
	return result
}
