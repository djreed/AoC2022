package main

import "testing"

func Test_calculateScore(t *testing.T) {
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
			if got := calculateScore(tt.args.choice, tt.args.opponent); got != tt.want {
				t.Errorf("calculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
