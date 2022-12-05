package main

import (
	"bufio"
	"fmt"
	"os"
)

type Backpack struct {
	contents string
	c1, c2   string
}

const (
	GROUP_SIZE = 3
)

type ElfGroup [GROUP_SIZE]Backpack

func main() {
	filename := "input"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elfBackpacks []Backpack
	for fileScanner.Scan() {
		line := fileScanner.Text()
		c1, c2 := splitBackpackString(line)
		elfBackpacks = append(elfBackpacks, Backpack{line, c1, c2})
	}

	fmt.Printf("Counted %d backpacks\n", len(elfBackpacks))

	var sharedItems []string
	for _, backpack := range elfBackpacks {
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

	var elfGroups []ElfGroup
	for i := 0; i < len(elfBackpacks); i += GROUP_SIZE {
		elfGroups = append(elfGroups, ElfGroup{elfBackpacks[i], elfBackpacks[i+1], elfBackpacks[i+2]})
	}

	fmt.Printf("Collected %d elf groups\n", len(elfGroups))

	var groupBadges []string
	for _, group := range elfGroups { // Group size is 3
		backpack1 := group[0]
		backpack2 := group[1]
		backpack3 := group[2]

		sharedItems1 := findSharedItems(backpack1.contents, backpack2.contents)
		sharedItems2 := findSharedItems(backpack2.contents, backpack3.contents)

		groupSharedItem := findSharedItem(sharedItems1, sharedItems2)
		groupBadges = append(groupBadges, groupSharedItem)
	}

	fmt.Printf("Collected %d group badges\n", len(groupBadges))

	var groupBadgeScores []int
	for _, groupBadge := range groupBadges {
		score := getItemScore(groupBadge)
		groupBadgeScores = append(groupBadgeScores, score)
	}

	fmt.Printf("Sum total of group badges: %d\n", Sum(groupBadgeScores))

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

func findSharedItems(c1, c2 string) string {
	var shared string
	for _, c1Item := range c1 {
		for _, c2Item := range c2 {
			if c1Item == c2Item {
				shared = shared + string(c1Item)
			}
		}
	}
	return shared
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
