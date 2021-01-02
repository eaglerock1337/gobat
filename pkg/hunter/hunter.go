/*
Package hunter is responsible for the primary logic of determining ideal Battleship
gameplay. Utilizing the basic structs and types implemented in the board module,
hunter implements its own structs and types for the seek-and-destroy process,
relying mostly on the built-in input validation and error checkintg from the board
module.

The PieceData type is responsible for maintaining the lists for all potential
ship placements based on a given ship's size (from 2 to 5 spaces). The HeatMap
type represents a simple 10x10 board of integers (similar to board.Board), but has
built-in methods for parsing the PieceData type and populating the heat map
accordingly.

The Hunter module ties all of this together by creating a larger struct with all
necessary variables needed to keep track of Battleship gameplay, including a board,
heatmap, lists of data, and other variables such as the amount of turns played.
It implements three major methods to do the following:
- Seek ships by finding the hottest potential squares in the heat map
- Destroy ships that have been found by shooting around known squares
- Take turns by accepting new data about the board and updating the board and piece data

I am thinking that the hitstack requires its own member methods and better error checking
and data handling, but that can be added later on.
*/
package hunter

import (
	"errors"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// These two variables allow for conversion of each square status to
// the status string and vice-versa. This allows for statuses to be stored
// as integers for faster lookup and comparison.
var (
	values = map[string]int{
		"Empty": 0, "Miss": 1, "Destroyer": 2, "Submarine": 3,
		"Cruiser": 4, "Battleship": 5, "Carrier": 6, "Hit": 7,
	}

	status = [8]string{
		"Empty", "Miss", "Destroyer", "Submarine",
		"Cruiser", "Battleship", "Carrier", "Hit",
	}
)

// The four directions (up, down, left, and right) for finding a ship in the hitmap
var directions = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

// Hunter is a struct that holds all data necessary to determine
// the optimal gameplay of Battleship.
type Hunter struct {
	Turns    int                // How many turns the Hunter has used
	Ships    []board.Ship       // The list of active unsunk ships
	Data     map[int]*PieceData // The list of possible ship positiions by size
	Board    board.Board        // The Battleship board with known data
	HeatMap  HeatMap            // The heat map popupated from the existing piece data
	SeekMode bool               // Whether the hunter is in Seek or Destroy mode
	Shots    []board.Square     // The current turn's list of best squares to play
	HitStack []board.Square     // The current number of outstanding hits
}

// NewHunter initializes a Hunter struct with the full list of ships,
// all possible ship locations, an empty board, and a heat map.
func NewHunter() Hunter {
	var newHunter Hunter
	newHunter.Ships = board.ShipTypes()
	newHunter.SeekMode = true

	for _, ship := range newHunter.Ships {
		if ship.GetType() != "Submarine" {
			data := GenPieceData(ship)
			newHunter.Data[ship.GetLength()] = &data
		}
	}

	return newHunter
}

// DeleteShip removes a ship from the list of active ships.
func (h *Hunter) DeleteShip(s board.Ship) error {
	for i, ship := range h.Ships {
		if ship.GetType() == s.GetType() {
			h.Ships[i] = h.Ships[len(h.Ships)-1]
			h.Ships = h.Ships[:len(h.Ships)-1]
			return nil
		}
	}
	return errors.New("Ship not found")
}

// GetValidLengths returns a slice of integers for all active ship types.
func (h Hunter) GetValidLengths() []int {
	var lengths []int
	foundThrees := false
	for _, ship := range h.Ships {
		length := ship.GetLength()
		if length == 3 {
			if foundThrees {
				continue
			}
			foundThrees = true
		}
		lengths = append(lengths, length)
	}

	return lengths
}

// AddHitStack adds a given Square to the hit stack.
// This probably requires error checking to ensure duplicates don't enter the stack.
func (h *Hunter) AddHitStack(s board.Square) {
	h.HitStack = append(h.HitStack, s)
}

// DelHitStack removes a given Square from the hit stack.
// This probably should return an error if it can't find the square.
func (h *Hunter) DelHitStack(s board.Square) {
	for i, square := range h.HitStack {
		if square.Letter == s.Letter && square.Number == s.Number {
			length := len(h.HitStack) - 1
			h.HitStack[i] = h.HitStack[length]
			h.HitStack = h.HitStack[:length]
			return
		}
	}
}

// InHitStack checks the hit stack if a given Square is present and returns a boolean.
func (h Hunter) InHitStack(s board.Square) bool {
	for _, square := range h.HitStack {
		if square.Letter == s.Letter && square.Number == s.Number {
			return true
		}
	}

	return false
}

// SearchPiece searches the PieceData for the given ship for all
// possible orientations, then intersect with the current hit stack.
// If the function succeeds in retrieving one result, it will return
// the piece with the location of the ship. Otherwise, the function
// will return with an error.
func (h Hunter) SearchPiece(sq board.Square, sh board.Ship) (board.Piece, error) {
	var hits []board.Piece
	length := sh.GetLength()
	for _, direction := range directions {
		let, num := direction[0], direction[1]

		// Check if the piece is in the stack
		nextSquare, err := board.SquareByValue(sq.Letter+let, sq.Number+num)
		if err == nil {
			continue
		}
		lastSquare, err := board.SquareByValue(sq.Letter+let*(length-1), sq.Number+num*(length-1))
		if err != nil {
			continue
		}
		for i := 1; i < length; i++ {
			square, _ := board.SquareByValue(sq.Letter+let*(i), sq.Number+num*(i))
			if !h.InHitStack(square) {
				continue
			}
		}

		// Create the piece to add to the list of hits
		var startSquare board.Square
		if let < 0 || num < 0 {
			startSquare = lastSquare
		} else {
			startSquare = nextSquare
		}

		var horizontal bool
		if let != 0 {
			horizontal = true
		} else {
			horizontal = false
		}

		piece, _ := board.NewPiece(sh, startSquare, horizontal)
		hits = append(hits, piece)
	}

	if len(hits) == 0 {
		return board.Piece{}, errors.New("No valid piece found in hit stack")
	}
	if len(hits) > 1 {
		return board.Piece{}, errors.New("Duplicate pieces found, algorithm failed")
	}
	return hits[0], nil
}

// SinkShip will use the active hit stack, the sinking square, and the
// type of ship sunk to find the exact location of the ship, update the
// board and piece data, as well as delete the ship from the ship list.
func (h *Hunter) SinkShip(sq board.Square, sh board.Ship) error {
	piece, err := h.SearchPiece(sq, sh)
	if err != nil {
		return errors.New("SinkShip failed due to SearchPiece not finding a piece")
	}

	for _, square := range piece.Coords {
		h.DelHitStack(square)
	}

	for _, length := range h.GetValidLengths() {
		h.Data[length].DeletePiece(piece)
	}

	err = h.DeleteShip(sh)
	if err != nil {
		return errors.New("SinkShip failed due to DeleteShip returning an error")
	}

	h.Board.SetPiece(piece)
	return nil
}

// Refresh will refresh the HeatMap based on the updated piece data and
// ship data.
func (h *Hunter) Refresh() {
	h.HeatMap.Initialize()

	for _, ship := range h.Ships {
		h.HeatMap.PopulateMap(*h.Data[ship.GetLength()], false)
	}
}

// Seek is the main hunting routine where the HeatMap is populated with
// all possible ship positions from the PieceData, and the top positions
// are populated in the Shots slice.
func (h *Hunter) Seek() {
	var top []board.Square
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			score := h.HeatMap[i][j]
			// Only try to add the value if it actually registered any hits
			if score > 0 {
				square, _ := board.SquareByValue(i, j)
				topLength := len(top)

				// Only add if the score is high enough or if the list isn't full yet
				if topLength < 5 || score > h.HeatMap.GetSquare(top[topLength-1]) {
					target := topLength
					if topLength > 0 {
						for k := topLength - 1; k >= 0; k-- {
							if score >= h.HeatMap.GetSquare(top[k]) {
								target = k
								break
							}
						}
					}

					top = append(top, square) // Add to empty array or make space
					if topLength > 0 {
						copy(top[target+1:topLength-1], top[target:])
						top[target] = square
					}
				}
			}
		}
	}

	h.Shots = top
}

// Destroy is the routine for sinking a ship that has been detected. Based
// on the squares in the HitStack, all available adjacent squares are
// checked in the HeatMap and ranked by total occurrences.
func (h *Hunter) Destroy() {

}

// Turn processes a single turn in the simulator based on the given
// square and result. The data is pruned, heatmap updated, and ideal moves
// given based on the mode the Hunter is currently in.
func (h *Hunter) Turn(s board.Square, result string) error {

}
