package hunter

import (
	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// HeatMap is a struct for holding heatmap data of a given Battleship board.
type HeatMap [10][10]int

// Initialize will zero out all values in the heatmap for reuse.
func (h *HeatMap) Initialize() {
	for i := range h {
		for j := range h[i] {
			h[i][j] = 0
		}
	}
}

// AddSquare will add one to the heatmap for the given Square.
func (h *HeatMap) AddSquare(s board.Square) {
	h[s.Letter][s.Number]++
}

// PopulateMap will add PieceData to the heatmap, optionally
// purging the existing data based on the initialize boolean.
func (h *HeatMap) PopulateMap(p PieceData, initialize bool) {
	if initialize {
		h.Initialize()
	}

	for _, piece := range p {
		for _, square := range piece.Coords {
			h.AddSquare(square)
		}
	}
}

// GetSquare will return the value of the given Square in the heatmap.
func (h HeatMap) GetSquare(s board.Square) int {
	return h[s.Letter][s.Number]
}
