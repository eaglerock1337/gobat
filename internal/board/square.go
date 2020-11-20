package board

import (
	"errors"
	"fmt"
)

// Square is a struct for holding a coordinate on the board
type Square struct {
	Letter int
	Number int
}

// StringByValue creates a Square by letter and number integers
func StringByValue(let int, num int) (Square, error) {
	if let < 0 || let > 9 || num < 0 || num > 9 {
		return Square{}, errors.New("String coordinates out of bounds")
	}
	return Square{let, num}, nil
}

// StringByString creates a Square by a string of the coordinates
func StringByString(coords string) (Square, error) {
	// This might not work every time
	return Square{0, 0}, nil
}

// PrintLetter returns the column (letter) as a string
func (s Square) PrintLetter() string {
	return string(rune('A' - 0 + s.Letter))
}

// PrintNumber returns the row (number) as a string
func (s Square) PrintNumber() string {
	return fmt.Sprint(s.Number + 1)
}

// PrintSquare returns the column and row as a two-letter string
func (s Square) PrintSquare() string {
	return s.PrintLetter() + s.PrintNumber()
}
