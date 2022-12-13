package main

import (
	"reflect"
	"testing"
)

func Test_splitElfPair(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name            string
		args            args
		wantAssignment1 ElfAssignment
		wantAssignment2 ElfAssignment
	}{
		{"example data", args{"2-4,5-6"}, ElfAssignment{2, 4}, ElfAssignment{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAssignment1, gotAssignment2 := splitElfPair(tt.args.line)
			if !reflect.DeepEqual(gotAssignment1, tt.wantAssignment1) {
				t.Errorf("splitElfPair() gotAssignment1 = %v, want %v", gotAssignment1, tt.wantAssignment1)
			}
			if !reflect.DeepEqual(gotAssignment2, tt.wantAssignment2) {
				t.Errorf("splitElfPair() gotAssignment2 = %v, want %v", gotAssignment2, tt.wantAssignment2)
			}
		})
	}
}

func Test_assignmentEncompassesAnother(t *testing.T) {
	type args struct {
		comparator ElfAssignment
		target     ElfAssignment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fully inclusive", args{ElfAssignment{1, 5}, ElfAssignment{2, 4}}, true},
		{"fully exclusive lower", args{ElfAssignment{1, 2}, ElfAssignment{3, 4}}, false},
		{"fully exclusive higher", args{ElfAssignment{3, 4}, ElfAssignment{1, 2}}, false},
		{"comparator start works only", args{ElfAssignment{1, 5}, ElfAssignment{2, 6}}, false},
		{"comparator end works only", args{ElfAssignment{3, 5}, ElfAssignment{2, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignmentEncompassesAnother(tt.args.comparator, tt.args.target); got != tt.want {
				t.Errorf("assignmentEncompassesAnother() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignmentOverlapsAnother(t *testing.T) {
	type args struct {
		comparator ElfAssignment
		target     ElfAssignment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fully inclusive", args{ElfAssignment{1, 5}, ElfAssignment{2, 4}}, true},
		{"fully exclusive lower", args{ElfAssignment{1, 2}, ElfAssignment{3, 4}}, false},
		{"fully exclusive higher", args{ElfAssignment{3, 4}, ElfAssignment{1, 2}}, false},
		{"comparator start works only", args{ElfAssignment{1, 5}, ElfAssignment{2, 6}}, true},
		{"comparator end works only", args{ElfAssignment{3, 5}, ElfAssignment{2, 4}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignmentOverlapsAnother(tt.args.comparator, tt.args.target); got != tt.want {
				t.Errorf("assignmentOverlapsAnother() = %v, want %v", got, tt.want)
			}
		})
	}
}
