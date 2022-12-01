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
	fmt.Printf("Top %d Elves are carrying a total of %d calories\n", requestedCount, Sum(sortedCalorieSumList[len(sortedCalorieSumList)-requestedCount:]))
}
