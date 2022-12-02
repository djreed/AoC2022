package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filename := "input"

	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elfCalorieList [][]int
	var elfCalories []int

	// Reach each line of the file, add it to a given elf's contents, wrap up that elf's contents if you encounter a newline
	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		lineAsInt, atoiErr := strconv.Atoi(lineStr)
		if atoiErr != nil { // On newline, append sum and reset elf calorie counter
			elfCalorieList = append(elfCalorieList, elfCalories)
			elfCalories = nil
		}
		elfCalories = append(elfCalories, lineAsInt)
	}
	// End of file, add whatever you've got left to the list
	elfCalorieList = append(elfCalorieList, elfCalories)
	readFile.Close()

	// Get a list of elf calorie sums
	var calorieSumList []int
	for _, elfCalories := range elfCalorieList {
		calorieSumList = append(calorieSumList, Sum(elfCalories))
	}

	// Get the max sum of values
	max := Max(calorieSumList)
	fmt.Printf("Highest Sum: %d\n", max)

	// Highest-to-Lowest sorted list
	sortedCalorieSumList := MergeSort(calorieSumList)
	requestedCount := 3

	fmt.Printf("Top %d highest calorie counts: %v\n", requestedCount, sortedCalorieSumList[len(sortedCalorieSumList)-requestedCount:])

	partialSum := Sum(sortedCalorieSumList[len(sortedCalorieSumList)-requestedCount:])
	fmt.Printf("Top %d Elves are carrying a total of %d calories\n", requestedCount, partialSum)
}

func Max(list []int) int {
	switch len(list) {
	case 0:
		return 0
	default:
		max := list[0]
		for _, val := range list[0:] {
			if max < val {
				max = val
			}
		}
		return max
	}
}

func Sum(list []int) int {
	result := 0
	for _, v := range list {
		result += v
	}
	return result
}

func MergeSort(list []int) []int {
	listLen := len(list)
	switch listLen {
	case 0:
		return nil
	case 1:
		return list
	default:
		halfLen := listLen / 2
		leftSide := list[0:halfLen]
		rightSide := list[halfLen:]

		mergedLeft := MergeSort(leftSide)
		mergedRight := MergeSort(rightSide)

		return merge(mergedLeft, mergedRight)
	}
}

// REQUIREMENT: la and lb are internally-sorted in ascending order
func merge(la, lb []int) []int {
	var mergedList []int

	for ia, ib := 0, 0; ia <= len(la) || ia <= len(lb); {
		if ia == len(la) {
			mergedList = append(mergedList, lb[ib:]...)
			break
		} else if ib == len(lb) {
			mergedList = append(mergedList, la[ia:]...)
			break
		} else {
			if la[ia] > lb[ib] {
				mergedList = append(mergedList, lb[ib])
				ib++
			} else {
				mergedList = append(mergedList, la[ia])
				ia++
			}
		}
	}

	return mergedList // Dead line, needed for compilation
}
