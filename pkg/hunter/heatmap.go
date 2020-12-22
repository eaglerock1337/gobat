package hunter

import (
	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// HeatMap is a slice of pieces that contains all possible placements for
// a piece with a given length. Each entry is stored as a slice of Squares.
type HeatMap [][]board.Square
