package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This is probably a bad idea, but I'm doing it anyway
var columns = map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
	"F": 6, "G": 7, "H": 8, "I": 9, "J": 10}

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
	chars := strings.Split(coords, "")
	if len(chars) > 2 {
		return Square{}, errors.New("Coordinate string has more than two characters")
	}

	if let, found := columns[chars[0]]; found {
		num, _ := strconv.Atoi(chars[1])
		if num < 0 || num > 9 {
			return Square{}, errors.New("String number coordinate out of bounds")
		}

		return Square{let, num}, nil
	}

	return Square{}, errors.New("String letter coordinate out of bounds")
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
