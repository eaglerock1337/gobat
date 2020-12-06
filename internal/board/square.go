package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This is probably a bad idea, but I'm doing it anyway
var columns = map[string]int{
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4,
	"F": 5, "G": 6, "H": 7, "I": 8, "J": 9,
}

// Square is a struct for holding a coordinate on the board
type Square struct {
	Letter int
	Number int
}

// SquareByValue creates a Square by letter and number integers
func SquareByValue(let int, num int) (Square, error) {
	if let < 0 || let > 9 || num < 0 || num > 9 {
		return Square{}, errors.New("String coordinates out of bounds")
	}
	return Square{let, num}, nil
}

// SquareByString creates a Square by a string of the coordinates
func SquareByString(coords string) (Square, error) {
	chars := strings.Split(coords, "")
	var letstr, numstr string

	if len(chars) < 2 || len(chars) > 3 {
		return Square{}, errors.New("Coordinate string is improperly sized")
	} else if len(chars) == 3 {
		letstr = chars[0]
		numstr = chars[1] + chars[2]
	} else {
		letstr = chars[0]
		numstr = chars[1]
	}

	if let, found := columns[strings.ToUpper(letstr)]; found {
		num, err := strconv.Atoi(numstr)
		if err != nil {
			return Square{}, errors.New("String number could not convert to integer")
		} else if num < 1 || num > 10 {
			return Square{}, errors.New("String number coordinate out of bounds")
		}

		return Square{let, num - 1}, nil
	}

	return Square{}, errors.New("String letter coordinate not found")
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
