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
*/
package hunter

import (
	"errors"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// Hunter is a struct that holds all data necessary to determine
// the optimal gameplay of Battleship.
type Hunter struct {
	Turns    int               // How many turns the Hunter has used
	Ships    []board.Ship      // The list of active unsunk ships
	Data     map[int]PieceData // The list of possible ship positiions by size
	Board    board.Board       // The Battleship board with known data
	HeatMap  HeatMap           // The heat map popupated from the existing piece data
	SeekMode bool              // Whether the hunter is in Seek or Destroy mode
	Shots    []board.Square    // The current turn's list of best squares to play
}

// NewHunter initializes a Hunter struct with the full list of ships,
// all possible ship locations, an empty board, and a heat map.
func NewHunter() Hunter {
	var newHunter Hunter
	newHunter.Ships = board.ShipTypes()
	newHunter.SeekMode = true

	for _, ship := range newHunter.Ships {
		if ship.GetString() != "Submarine" {
			newHunter.Data[ship.GetLength()] = GenPieceData(ship)
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

// Seek is the main hunting routine where the HeatMap is populated with
// all possible ship positions from the PieceData, and the top positions
// are populated in the Shots slice.
func (h *Hunter) Seek() {

}
