package board

import "errors"

var ships = map[string]int{
	"Carrier":    5,
	"Battleship": 4,
	"Cruiser":    3,
	"Submarine":  3,
	"Destroyer":  2,
}

// Ship is a string of the ship type with extra methods
type Ship string

// Piece is a struct for defining a piece and its position
type Piece struct {
	Type   Ship
	Coords []Square
}

// NewShip will return a Ship type after input validation
func NewShip(shipType string) (Ship, error) {
	if _, ok := ships[shipType]; ok {
		return Ship(shipType), nil
	}
	return Ship(""), errors.New("A valid Ship type was not given")
}

// Type will return the type of ship as a string
func (s Ship) Type() string {
	return string(s)
}

// Length will return the length of the ship as an integer
func (s Ship) Length() int {
	return ships[string(s)]
}

// ShipTypes will return a slice of all 5 Ship types
func ShipTypes() []Ship {
	shipTypes := make([]Ship, 0, 5)
	for ship := range ships {
		shipTypes = append(shipTypes, Ship(ship))
	}
	return shipTypes
}

// NewPiece defines a Piece by the following variables:
func NewPiece(shipType Ship, startSquare Square, horizontal bool) (Piece, error) {
	// TODO: do piece on board validation here
	var newPiece Piece
	newPiece.Type = shipType
	newPiece.Coords = make([]Square, shipType.Length())
	return newPiece, nil

	// return Piece{}, errors.New("Something went wrong")
}
