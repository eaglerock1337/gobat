package hunter

import (
	"errors"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// HitStack is a type that provides a stack of squares that indicate hits
// in the current game that are not yet associated with a sunken ship.
type HitStack []board.Square

// Push adds a given Square to the hit stack, returning an error if
// a duplicate entry is already present.
func (h *HitStack) Push(s board.Square) error {
	for _, entry := range *h {
		if entry.Letter == s.Letter && entry.Number == s.Number {
			return errors.New("Push failed, duplicate entry found in hit stack")
		}
	}

	*h = append(*h, s)
	return nil
}

// Pop removes a given Square from the hit stack, returning an error
// if the square is not found.
func (h *HitStack) Pop(s board.Square) error {
	for i, square := range *h {
		if square.Letter == s.Letter && square.Number == s.Number {
			length := len(*h) - 1
			(*h)[i] = (*h)[length]
			*h = (*h)[:length]
			return nil
		}
	}
	return errors.New("Pop failed, square not found in hit stack")
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
