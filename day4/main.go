package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfAssignment struct {
	start, end int
}

type ElfPair struct {
	line   string
	e1, e2 ElfAssignment
}

func main() {
	filename := "input"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elfPairs []ElfPair
	for fileScanner.Scan() {
		line := fileScanner.Text()
		c1, c2 := splitElfPair(line)
		elfPairs = append(elfPairs, ElfPair{line, c1, c2})
	}

	fmt.Printf("Counted %d pairs of Elf assignments\n", len(elfPairs))

	containedPairs := 0
	for _, pair := range elfPairs {
		if assignmentEncompassesAnother(pair.e1, pair.e2) || assignmentEncompassesAnother(pair.e2, pair.e1) {
			containedPairs++
		}
	}

	fmt.Printf("%d Elf pairs feature one assignment enclosing another\n", containedPairs)

	overlappingPairs := 0
	for _, pair := range elfPairs {
		if assignmentOverlapsAnother(pair.e1, pair.e2) {
			overlappingPairs++
		}
	}

	fmt.Printf("%d Elf pairs feature one assignment overlapping another\n", overlappingPairs)

}

// Parse elf pairs
// Ex: 2-5,15-90
func splitElfPair(line string) (assignment1, assignment2 ElfAssignment) {
	substrings := strings.Split(line, ",")
	assignment1String, assignment2String := substrings[0], substrings[1]

	assignment1Split := strings.Split(assignment1String, "-")
	assignment1.start, _ = strconv.Atoi(assignment1Split[0])
	assignment1.end, _ = strconv.Atoi(assignment1Split[1])

	assignment2Split := strings.Split(assignment2String, "-")
	assignment2.start, _ = strconv.Atoi(assignment2Split[0])
	assignment2.end, _ = strconv.Atoi(assignment2Split[1])

	return // assignment1, assignment2
}

// Return whether the first range encloses the second
func assignmentEncompassesAnother(comparator, target ElfAssignment) bool {
	return comparator.start <= target.start && comparator.end >= target.end
}

// Return whether the two ranges overlap
func assignmentOverlapsAnother(comparator, target ElfAssignment) bool {
	encompassesOne := assignmentEncompassesAnother(comparator, target)
	encompassesTwo := assignmentEncompassesAnother(target, comparator)

	// cs --- ce
	//    ts --- te
	leftOverlap := (comparator.start <= target.start && comparator.start <= target.end) &&
		(target.start >= comparator.start && target.start <= comparator.end)

	// ts --- te
	//    cs --- ce
	rightOverlap := (target.start <= comparator.start && target.start <= comparator.end) &&
		(comparator.start >= target.start && comparator.start <= target.end)

	return encompassesOne || encompassesTwo || leftOverlap || rightOverlap
}
