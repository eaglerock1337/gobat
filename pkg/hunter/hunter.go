package hunter

import (
	"errors"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// Hunter is a struct that holds all data necessary to determine
// the optimal gameplay of Battleship. Currently, the data will
// duplicate the info for the Cruiser and the Submarine, but this
// can be optimized later.
type Hunter struct {
	Ships    []board.Ship
	Data     map[int]PieceData
	Board    board.Board
	HeatMap  HeatMap
	SeekMode bool
	Shots    []board.Square
}

// NewHunter initializes a Hunter struct with the full list of ships,
// all possible ship locations, an empty board, and a heat map.
func NewHunter() Hunter {
	var newHunter Hunter
	newHunter.Ships = board.ShipTypes()
	newHunter.SeekMode = false

	for _, ship := range newHunter.Ships {
		newHunter.Data[ship.GetLength()] = GenPieceData(ship)
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

// Hunt is the main hunting routine where the HeatMap is populated with
// all possible ship positions from the PieceData, and the top positions
// are populated in the Shots slice.
func (h *Hunter) Hunt() {

}
