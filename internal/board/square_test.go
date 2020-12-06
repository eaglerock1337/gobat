package board

import (
	"strings"
	"testing"
)

var exampleValues = [5][2]int{
	{0, 0},
	{2, 2},
	{5, 2},
	{9, 6},
	{6, 9},
}

var exampleStrings = [5]string{"A1", "c3", "F3", "J7", "g10"}

var examples = [5]Square{
	{0, 0},
	{2, 2},
	{5, 2},
	{9, 6},
	{6, 9},
}

var badValues = [5][2]int{
	{-1, 0},
	{0, 10},
	{11, 5},
	{-1, -1},
	{12, -12},
}

var badStrings = [5]string{"Z1", "3A", "AA10", "AA1", "B11"}

func TestSquareByValue(t *testing.T) {
	for i, input := range exampleValues {
		answer, err := SquareByValue(input[0], input[1])

		if err != nil {
			t.Errorf("SquareByValue returned an error: %v", err)
		} else if answer != examples[i] {
			t.Errorf("SquareByValue function was incorrect, got: %v, want: %v", answer, examples[i])
		}
	}
}

func TestBadSquareByValue(t *testing.T) {
	for _, input := range badValues {
		answer, err := SquareByValue(input[0], input[1])

		if err == nil {
			t.Errorf("SquareByValue did not error as expected with %v, returned Square: %v", input, answer)
		}
	}
}

func TestSquareByString(t *testing.T) {
	for i, input := range exampleStrings {
		answer, err := SquareByString(input)

		if err != nil {
			t.Errorf("SquareByString returned an error: %v", err)
		} else if answer != examples[i] {
			t.Errorf("SquareByString function was incorrect, got: %v, want: %v", answer, examples[i])
		}
	}
}

func TestBadSquareByString(t *testing.T) {
	for _, input := range badStrings {
		answer, err := SquareByString(input)

		if err == nil {
			t.Errorf("SquareByString did not error as expected with %v, returned Square: %v", input, answer)
		}
	}
}

func TestPrintLetter(t *testing.T) {
	results := [5]string{"A", "C", "F", "J", "G"}

	for i := 0; i < 5; i++ {
		answer := examples[i].PrintLetter()
		if answer != results[i] {
			t.Errorf("Letter function was incorrect, got: %v, want: %v", answer, results[i])
		}
	}
}

func TestPrintNumber(t *testing.T) {
	results := [5]string{"1", "3", "3", "7", "10"}

	for i := 0; i < 5; i++ {
		answer := examples[i].PrintNumber()
		if answer != results[i] {
			t.Errorf("Letter function was incorrect, got: %v, want: %v", answer, results[i])
		}
	}
}

func TestPrintSquare(t *testing.T) {
	for i := 0; i < 4; i++ {
		answer := examples[i].PrintSquare()
		if answer != strings.ToUpper(exampleStrings[i]) {
			t.Errorf("Square function was incorrect, got: %v, want: %v", answer, exampleStrings[i])
		}
	}
}
