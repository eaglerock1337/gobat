package hunter

import (
	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// HitStack is a type that provides a stack of squares that indicate hits
// in the current game that are not yet associated with a sunken ship.
type HitStack []board.Square

// Push adds a given Square to the hit stack.
// This probably requires error checking to ensure duplicates don't enter the stack.
func (h *HitStack) Push(s board.Square) {
	*h = append(*h, s)
}

// Pop removes a given Square from the hit stack.
// This probably should return an error if it can't find the square.
func (h *HitStack) Pop(s board.Square) {
	for i, square := range *h {
		if square.Letter == s.Letter && square.Number == s.Number {
			length := len(*h) - 1
			(*h)[i] = (*h)[length]
			*h = (*h)[:length]
			return
		}
	}
}

// InStack checks the hit stack if a given Square is present and returns a boolean.
func (h HitStack) InStack(s board.Square) bool {
	for _, square := range h {
		if square.Letter == s.Letter && square.Number == s.Number {
			return true
		}
	}

	return false
}
