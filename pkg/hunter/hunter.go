/*
Package hunter is responsible for the primary logic of determining ideal Battleship
gameplay. Utilizing the basic structs and types implemented in the board module,
hunter implements its own structs and types for the seek-and-destroy process,
relying mostly on the built-in input validation and error checking from the board
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
	"fmt"

	"github.com/eaglerock1337/gobat/pkg/board"
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

// The four directions (up, down, left, and right) for finding adjacent squares
var directions = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

// Hunter is a struct that holds all data necessary to determine
// the optimal gameplay of Battleship.
type Hunter struct {
	Turns    int                // How many turns the Hunter has used
	Ships    []board.Ship       // The list of active unsunk ships
	Data     map[int]*PieceData // The list of possible ship positions by size
	Board    board.Board        // The Battleship board with known data
	HeatMap  HeatMap            // The heat map populated from the existing piece data
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
	newHunter.Shots = make([]board.Square, 0, 5)
	newHunter.Data = make(map[int]*PieceData)

	for _, ship := range newHunter.Ships {
		if ship.GetType() != "Submarine" {
			data := GenPieceData(ship)
			newHunter.Data[ship.GetLength()] = &data
		}
	}

	newHunter.Refresh()
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
	return errors.New("unable to find Ship to delete")
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

// InShots checks if the current square is in the shot list and returns a boolean.
func (h Hunter) InShots(s board.Square) bool {
	for _, square := range h.Shots {
		if square.Letter == s.Letter && square.Number == s.Number {
			return true
		}
	}

	return false
}

// Refresh will refresh the HeatMap based on the updated piece data and
// ship data.
func (h *Hunter) Refresh() {
	h.HeatMap.Initialize()

	for _, ship := range h.Ships {
		h.HeatMap.PopulateMap(*h.Data[ship.GetLength()], false)
	}
}

// AddShot will attempt to add the given square to the Shots array, which
// will only get accepted if in the top 5 Shots.
func (h *Hunter) AddShot(s board.Square) {
	score := h.HeatMap[s.Letter][s.Number]
	// Only try to add the value if not in the hitstack and if it registered a score
	if score > 0 && !h.InHitStack(s) {
		length := len(h.Shots)

		// Only add if the score is high enough or if the list isn't full yet
		if length < 5 || score > h.HeatMap.GetSquare(h.Shots[length-1]) {
			target := length - 1
			if length > 0 {
				for k := 0; k < length-1; k++ {
					if score >= h.HeatMap.GetSquare(h.Shots[k]) {
						target = k
						break
					}
				}
			}

			if length != 5 {
				h.Shots = append(h.Shots, s) // Add to empty array or make space
			}

			if length > 0 {
				copy(h.Shots[target+1:], h.Shots[target:length])
				h.Shots[target] = s
			}
		}
	}
}

// ClearShots will empty out the current shot list.
func (h *Hunter) ClearShots() {
	h.Shots = make([]board.Square, 0, 5)
}

// SearchPiece searches the PieceData for the given ship for all
// possible orientations, then intersect with the current hit stack.
// If the function succeeds in retrieving one result, it will return
// the piece with the location of the ship. Otherwise, the function
// will return with an error.
func (h Hunter) SearchPiece(sq board.Square, sh board.Ship) (board.Piece, error) {
	var hits []board.Piece
	length := sh.GetLength()
Direction:
	for _, direction := range directions {
		let, num := direction[0], direction[1]

		// Check if the piece is in the stack
		_, err := board.SquareByValue(sq.Letter+let, sq.Number+num)
		if err != nil {
			continue
		}
		lastSquare, err := board.SquareByValue(sq.Letter+let*(length-1), sq.Number+num*(length-1))
		if err != nil {
			continue
		}
		for i := 0; i < length; i++ {
			square, _ := board.SquareByValue(sq.Letter+let*(i), sq.Number+num*(i))

			if !h.InHitStack(square) {
				continue Direction
			}
		}

		// Create the piece to add to the list of hits
		startSquare := sq
		if lastSquare.Letter < sq.Letter || lastSquare.Number < sq.Number {
			startSquare = lastSquare
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
		return board.Piece{}, errors.New("no valid piece found in hit stack")
	}
	if len(hits) > 1 {
		return board.Piece{}, errors.New("duplicate pieces found, algorithm failed")
	}
	return hits[0], nil
}

// SinkShip will use the active hit stack, the sinking square, and the
// type of ship sunk to find the exact location of the ship, update the
// board and piece data, as well as delete the ship from the ship list.
func (h *Hunter) SinkShip(sq board.Square, sh board.Ship) error {
	piece, err := h.SearchPiece(sq, sh)
	if err != nil {
		return fmt.Errorf("SinkShip failed due to SearchPiece not finding a piece: %v", err)
	}

	err = h.DeleteShip(sh)
	if err != nil {
		return fmt.Errorf("SinkShip failed due to DeleteShip returning error: %v", err)
	}

	for _, square := range piece.Coords {
		h.DelHitStack(square)
	}

	for _, length := range h.GetValidLengths() {
		h.Data[length].DeletePiece(piece)
	}

	h.Board.SetPiece(piece)
	return nil
}

// Seek is the main hunting routine where the HeatMap is populated with
// all possible ship positions from the PieceData, and the top positions
// are populated in the Shots slice.
func (h *Hunter) Seek() {
	h.ClearShots()

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			square, _ := board.SquareByValue(i, j)
			h.AddShot(square)
		}
	}
}

// Destroy is the routine for sinking a ship that has been detected. Based
// on the squares in the HitStack, all available adjacent squares are
// checked in the HeatMap and ranked by total occurrences.
func (h *Hunter) Destroy() {
	h.ClearShots()

	for _, hit := range h.HitStack {
		for _, direction := range directions {
			let, num := direction[0], direction[1]
			square, err := board.SquareByValue(hit.Letter+let, hit.Number+num)
			if err == nil {
				for _, shot := range h.Shots {
					if shot.Letter == square.Letter && shot.Number == square.Number {
						continue
					}
				}
				h.AddShot(square)
			}
		}
	}
}

// Turn processes a single turn in the simulator based on the given
// square and result. The data is pruned, heatmap updated, and ideal moves
// given based on the mode the Hunter is currently in.
func (h *Hunter) Turn(s board.Square, result string) error {
	err := h.Board.SetString(s, result)
	if err != nil {
		return fmt.Errorf("Turn failed as the result was invalid: %v", err)
	}

	if h.Board.IsEmpty(s) {
		return errors.New("Turn failed as it was given an empty result")
	}

	origSeekMode := h.SeekMode // save state in case of SinkShip error below

	if h.Board.IsHit(s) {
		h.AddHitStack(s)
		h.SeekMode = false
	}

	if h.Board.IsSunk(s) {
		ship, _ := board.NewShip(result)
		err := h.SinkShip(s, ship)
		if err != nil { // undo changes from recording a hit before erroring out
			h.SeekMode = origSeekMode
			h.Board.SetString(s, "Empty")
			h.DelHitStack(s)
			return fmt.Errorf("Turn failed due to SinkShip error: %v", err)
		}
		h.SeekMode = len(h.HitStack) == 0
	}

	if h.Board.IsMiss(s) {
		for _, length := range h.GetValidLengths() {
			h.Data[length].DeleteSquare(s)
		}
	}

	h.Refresh()

	if h.SeekMode {
		h.Seek()
	} else {
		h.Destroy()
	}

	h.Turns++
	return nil
}
