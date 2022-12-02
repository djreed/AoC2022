package main

import (
	"testing"
)

func Test_calculateScoreChoices(t *testing.T) {
	type args struct {
		choice   string
		opponent string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"rock-rock", args{ROCK, ROCK}, 4},
		{"rock-paper", args{ROCK, PAPER}, 1},
		{"rock-scissors", args{ROCK, SCISSORS}, 7},
		{"paper-rock", args{PAPER, ROCK}, 8},
		{"paper-paper", args{PAPER, PAPER}, 5},
		{"paper-scissors", args{PAPER, SCISSORS}, 2},
		{"scissors-rock", args{SCISSORS, ROCK}, 3},
		{"scissors-paper", args{SCISSORS, PAPER}, 9},
		{"scissors-scissors", args{SCISSORS, SCISSORS}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScoreChoices(tt.args.choice, tt.args.opponent); got != tt.want {
				t.Errorf("calculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getChoiceForOutcome(t *testing.T) {
	type args struct {
		opponentChoice string
		outcome        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"rock - win", args{ROCK, WIN}, PAPER},
		{"rock - loss", args{ROCK, LOSS}, SCISSORS},
		{"rock - draw", args{ROCK, DRAW}, ROCK},

		{"paper - win", args{PAPER, WIN}, SCISSORS},
		{"paper - loss", args{PAPER, LOSS}, ROCK},
		{"paper - draw", args{PAPER, DRAW}, PAPER},

		{"scissors - win", args{SCISSORS, WIN}, ROCK},
		{"scissors - loss", args{SCISSORS, LOSS}, PAPER},
		{"scissors - draw", args{SCISSORS, DRAW}, SCISSORS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChoiceForOutcome(tt.args.opponentChoice, tt.args.outcome); got != tt.want {
				t.Errorf("getChoiceForOutcome() = %v, want %v", got, tt.want)
			}
		})
	}
}
