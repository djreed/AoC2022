package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"

	WIN  = 6
	DRAW = 3
)

var (
	SCORES = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	CODE_MAP = map[string]string{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSORS,
		"Z": SCISSORS,
	}
)

func main() {
	filename := "input"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	selfScore := 0

	// Reach each line of the file, add it to a given elf's contents, wrap up that elf's contents if you encounter a newline
	for fileScanner.Scan() {
		lineStr := fileScanner.Text()

		lineContents := strings.Split(lineStr, " ")

		selfCode := lineContents[1]
		opponentCode := lineContents[0]

		choice := CODE_MAP[selfCode]
		opponentChoice := CODE_MAP[opponentCode]

		roundScore := calculateScore(choice, opponentChoice)
		fmt.Printf("Score from %s (%s) <> %s (%s) = %d\n",
			selfCode, choice, opponentCode, opponentChoice, roundScore)

		selfScore += roundScore
	}

	fmt.Printf("Total Score: %d\n", selfScore)
}

func calculateScore(choice, opponent string) int {
	score := 0

	// Score value of what you chose
	score += SCORES[choice]

	// Same pick, draw game
	if choice == opponent {
		return score + DRAW
	}

	// All losses enumerated
	switch choice {
	case ROCK:
		if opponent == PAPER {
			return score
		}
	case PAPER:
		if opponent == SCISSORS {
			return score
		}
	case SCISSORS:
		if opponent == ROCK {
			return score
		}
	}

	// Not a draw, not a loss, you win!
	return score + WIN
}
