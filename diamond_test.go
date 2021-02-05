package main

import (
	"reflect"
	"testing"
)

func Test_alphabetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Test 01", input: "A", want: 0},
		{name: "Test 02", input: "B", want: 1},
		{name: "Test 03", input: "C", want: 2},
		{name: "Test 04", input: "D", want: 3},
		{name: "Test 05", input: "E", want: 4},
		{name: "Test 06", input: "F", want: 5},
		{name: "Test 07", input: "G", want: 6},
		{name: "Test 08", input: "H", want: 7},
		{name: "Test 09", input: "I", want: 8},
		{name: "Test 10", input: "J", want: 9},
		{name: "Test 11", input: "K", want: 10},
		{name: "Test 12", input: "L", want: 11},
		{name: "Test 13", input: "M", want: 12},
		{name: "Test 14", input: "N", want: 13},
		{name: "Test 15", input: "O", want: 14},
		{name: "Test 16", input: "P", want: 15},
		{name: "Test 17", input: "Q", want: 16},
		{name: "Test 18", input: "R", want: 17},
		{name: "Test 19", input: "S", want: 18},
		{name: "Test 20", input: "T", want: 19},
		{name: "Test 21", input: "U", want: 20},
		{name: "Test 22", input: "V", want: 21},
		{name: "Test 23", input: "W", want: 22},
		{name: "Test 24", input: "X", want: 23},
		{name: "Test 25", input: "Y", want: 24},
		{name: "Test 26", input: "Z", want: 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabetValue(tt.input); got != tt.want {
				t.Errorf("alphabetValue(tt.input) = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueAlphabet(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "Test 01", input: 0, want: "A"},
		{name: "Test 02", input: 1, want: "B"},
		{name: "Test 03", input: 2, want: "C"},
		{name: "Test 04", input: 3, want: "D"},
		{name: "Test 05", input: 4, want: "E"},
		{name: "Test 06", input: 5, want: "F"},
		{name: "Test 07", input: 6, want: "G"},
		{name: "Test 08", input: 7, want: "H"},
		{name: "Test 09", input: 8, want: "I"},
		{name: "Test 10", input: 9, want: "J"},
		{name: "Test 11", input: 10, want: "K"},
		{name: "Test 12", input: 11, want: "L"},
		{name: "Test 13", input: 12, want: "M"},
		{name: "Test 14", input: 13, want: "N"},
		{name: "Test 15", input: 14, want: "O"},
		{name: "Test 16", input: 15, want: "P"},
		{name: "Test 17", input: 16, want: "Q"},
		{name: "Test 18", input: 17, want: "R"},
		{name: "Test 19", input: 18, want: "S"},
		{name: "Test 20", input: 19, want: "T"},
		{name: "Test 21", input: 20, want: "U"},
		{name: "Test 22", input: 21, want: "V"},
		{name: "Test 23", input: 22, want: "W"},
		{name: "Test 24", input: 23, want: "X"},
		{name: "Test 25", input: 24, want: "Y"},
		{name: "Test 26", input: 25, want: "Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueAlphabet(tt.input); got != tt.want {
				t.Errorf("valueAlphabet(tt.input) = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSpacing(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		want   string
		length int
	}{
		{name: "Test 01", input: 0, want: ``, length: 0},
		{name: "Test 02", input: 1, want: ` `, length: 1},
		{name: "Test 03", input: 2, want: `  `, length: 2},
		{name: "Test 04", input: 3, want: `   `, length: 3},
		{name: "Test 10", input: 10, want: `          `, length: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSpacing(tt.input); got != tt.want { //len(got) != tt.length;
				t.Errorf("stringSpacing(tt.input) = %v, want %v", len(got), tt.length)
			}
		})
	}
}
func Test_drawDiamond(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   []string
		length int
	}{
		{name: "Test 01", input: "A", want: []string{"A"}, length: 1},
		{name: "Test 02", input: "C", want: []string{"  A  ", " B B ", "C   C", " B B ", "  A  "}, length: 25},
		{name: "Test 03", input: "E", want: []string{"    A    ", "   B B   ", "  C   C  ", " D     D ", "E       E", " D     D ", "  C   C  ", "   B B   ", "    A    "}, length: 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawDiamond(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawDiamond(tt.input) = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcSpacing(t *testing.T) {
	tests := []struct {
		name        string
		letterValue int
		loop        int
		want        int
	}{
		{name: "Test 01", letterValue: 0, loop: 0, want: 0},
		{name: "Test 02", letterValue: 2, loop: 0, want: 2},
		{name: "Test 03", letterValue: 4, loop: 0, want: 4},
		{name: "Test 04", letterValue: 4, loop: 1, want: 3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSpacing(tt.letterValue, tt.loop); got != tt.want {
				t.Errorf("calcSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcMidSpacing(t *testing.T) {
	tests := []struct {
		name string
		loop int
		want int
	}{
		{name: "Test 01", loop: 0, want: -1},
		{name: "Test 02", loop: 1, want: 1},
		{name: "Test 03", loop: 2, want: 3},
		{name: "Test 04", loop: 3, want: 5},
		{name: "Test 05", loop: 4, want: 7},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMidSpacing(tt.loop); got != tt.want {
				t.Errorf("calcMidSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcGridSize(t *testing.T) {
	tests := []struct {
		name        string
		letterValue int
		want        int
	}{
		{name: "Test 01", letterValue: 0, want: 1},
		{name: "Test 02", letterValue: 1, want: 3},
		{name: "Test 03", letterValue: 2, want: 5},
		{name: "Test 04", letterValue: 3, want: 7},
		{name: "Test 05", letterValue: 4, want: 9},
		{name: "Test 06", letterValue: 5, want: 11},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcGridSize(tt.letterValue); got != tt.want {
				t.Errorf("calcGridSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_innerDrawLoop(t *testing.T) {
	tests := []struct {
		name        string
		loop        int
		LetterValue int
		want        string
	}{
		{name: "Test 01", loop: 0, LetterValue: 0, want: "A"},
		{name: "Test 02", loop: 0, LetterValue: 2, want: "  A  "},
		{name: "Test 03", loop: 0, LetterValue: 4, want: "    A    "},
		{name: "Test 04", loop: 2, LetterValue: 4, want: "  C   C  "},
		{name: "Test 05", loop: 0, LetterValue: 5, want: "     A     "},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := innerDrawLoop(tt.loop, tt.LetterValue); got != tt.want {
				t.Errorf("innerDrawLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
