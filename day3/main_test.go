package main

import (
	"testing"
)

func Test_splitBackpackStrings(t *testing.T) {
	type args struct {
		contents string
	}
	tests := []struct {
		name   string
		args   args
		wantC1 string
		wantC2 string
	}{
		{"empty string", args{""}, "", ""},
		{"length is one", args{"a"}, "", "a"},
		{"length is two", args{"ab"}, "a", "b"},
		{"example backpack", args{"vJrwpWtwJgWrhcsFMMfFFhFp"}, "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC1, gotC2 := splitBackpackString(tt.args.contents)
			if gotC1 != tt.wantC1 {
				t.Errorf("splitBackpackStrings() gotComp1 = %v, want %v", gotC1, tt.wantC1)
			}
			if gotC2 != tt.wantC2 {
				t.Errorf("splitBackpackStrings() gotComp2 = %v, want %v", gotC2, tt.wantC2)
			}
		})
	}
}

func Test_findSharedItem(t *testing.T) {
	type args struct {
		c1 string
		c2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty strings", args{"", ""}, ""},
		{"example backpack", args{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, "p"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSharedItem(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("findSharedItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getItemScore(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"letter a", args{"a"}, 1},
		{"letter z", args{"z"}, 26},
		{"letter A", args{"A"}, 27},
		{"letter Z", args{"Z"}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getItemScore(tt.args.item); got != tt.want {
				t.Errorf("getItemScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
