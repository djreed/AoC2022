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

	WIN  = "win"
	DRAW = "draw"
	LOSS = "lose"

	WIN_SCORE  = 6
	DRAW_SCORE = 3
	LOSS_SCORE = 0

	WIN_CODE  = "X"
	DRAW_CODE = "Y"
	LOSS_CODE = "Z"
)

var (
	CHOICE_SCORES = map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}
	RESULT_SCORE = map[string]int{
		WIN:  6,
		DRAW: 3,
		LOSS: 0,
	}
	WINS_AGAINST = map[string]string{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}
	LOSES_TO = map[string]string{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}

	CODE_MAP_SCORE = map[string]string{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	CODE_MAP_OUTCOME = map[string]string{
		"X": LOSS,
		"Y": DRAW,
		"Z": WIN,
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

	for fileScanner.Scan() {
		lineStr := fileScanner.Text()

		lineContents := strings.Split(lineStr, " ")

		opponentCode := lineContents[0]
		resultCode := lineContents[1]

		opponentChoice := CODE_MAP_SCORE[opponentCode]
		resultOutcome := CODE_MAP_OUTCOME[resultCode]
		choice := getChoiceForOutcome(opponentChoice, resultOutcome)

		roundScore := calculateScoreChoices(choice, opponentChoice)
		fmt.Printf("Score from '%s %s' => %s<>%s => %d\n",
			opponentCode, resultCode,
			opponentChoice, choice,
			roundScore)

		selfScore += roundScore
	}

	fmt.Printf("Total Score: %d\n", selfScore)
}

// Choice + Opponent -> round score
func calculateScoreChoices(choice, opponent string) int {
	// Score value of what you chose
	score := CHOICE_SCORES[choice]

	switch choice {
	case WINS_AGAINST[opponent]:
		return score + RESULT_SCORE[WIN] // Win
	case LOSES_TO[opponent]:
		return score + RESULT_SCORE[LOSS] // Loss
	default:
		return score + RESULT_SCORE[DRAW] // Draw
	}
}

// Opponent + Outcome -> Choice to get that outcome
func getChoiceForOutcome(opponentChoice, outcome string) string {
	switch outcome {
	case WIN:
		return WINS_AGAINST[opponentChoice]
	case LOSS:
		return LOSES_TO[opponentChoice]
	default:
		return opponentChoice
	}
}
