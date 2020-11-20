package board

import "fmt"

// Square is a struct for holding a coordinate on the board
type Square struct {
	letter int
	number int
}

// Letter returns the column (letter) as a string
func (s Square) Letter() string {
	return string(rune('A' - 0 + s.letter))
}

// Number returns the row (number) as a string
func (s Square) Number() string {
	return fmt.Sprint(s.number + 1)
}

// Square returns the column and row as a two-letter string
func (s Square) Square() string {
	return s.Letter() + s.Number()
}
