package board

import "errors"

// I can optimize this later
var ships = map[string]int{
	"Carrier":    5,
	"Battleship": 4,
	"Cruiser":    3,
	"Submarine":  3,
	"Destroyer":  2,
}

// I can't decide if it's smart do this just to avoid computational
// complexity, or if I am just being horribly lazy
var shipNames = [5]string{"Carrier", "Battleship", "Cruiser", "Submarine", "Destroyer"}

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
	for _, ship := range shipNames {
		shipTypes = append(shipTypes, Ship(ship))
	}
	return shipTypes
}

// NewPiece defines a Piece by a ship type, a starting coordinate, and the
// direction (horizontal or vertical), and returns a Piece and error result
func NewPiece(shipType Ship, startSquare Square, horizontal bool) (Piece, error) {
	var newPiece Piece
	newPiece.Type = shipType
	newPiece.Coords = make([]Square, 0, shipType.Length())

	for i := 0; i < shipType.Length(); i++ {
		letter := startSquare.Letter
		number := startSquare.Number
		if horizontal {
			letter += i
		} else {
			number += i
		}

		square, error := SquareByValue(letter, number)
		if error != nil {
			return newPiece, errors.New("Ship location is out of bounds")
		}
		newPiece.Coords = append(newPiece.Coords, square)
	}

	return newPiece, nil
}
