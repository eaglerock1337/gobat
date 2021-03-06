package board

import "errors"

// ships provides the list of ships in Battleship and their lengths.
var ships = map[string]int{
	"Carrier":    5,
	"Battleship": 4,
	"Cruiser":    3,
	"Submarine":  3,
	"Destroyer":  2,
}

// shipNames provides the list of ships in Battleship in array form.
var shipNames = [5]string{"Carrier", "Battleship", "Cruiser", "Submarine", "Destroyer"}

// Ship is a string of the ship type with extra methods.
type Ship string

// Piece is a struct for defining a piece and its position.
type Piece struct {
	Type   Ship
	Coords []Square
}

// Ship creation functions

// NewShip will return a Ship type after input validation.
func NewShip(shipType string) (Ship, error) {
	if _, ok := ships[shipType]; ok {
		return Ship(shipType), nil
	}
	return Ship(""), errors.New("A valid Ship type was not given")
}

// ShipTypes will return a slice of all 5 Ship types.
func ShipTypes() []Ship {
	shipTypes := make([]Ship, 0, 5)
	for _, ship := range shipNames {
		shipTypes = append(shipTypes, Ship(ship))
	}
	return shipTypes
}

// Ship retrieval methods

// GetType will return the type of ship as a string.
func (s Ship) GetType() string {
	return string(s)
}

// GetLength will return the length of the ship as an integer.
func (s Ship) GetLength() int {
	return ships[string(s)]
}

// Piece creation function

// NewPiece defines a Piece by a ship type, a starting coordinate, and the
// direction (horizontal or vertical), and returns a Piece and error result.
func NewPiece(shipType Ship, startSquare Square, horizontal bool) (Piece, error) {
	var newPiece Piece
	newPiece.Type = shipType
	newPiece.Coords = make([]Square, 0, shipType.GetLength())

	for i := 0; i < shipType.GetLength(); i++ {
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

// Piece boolean methods

// InSquare is a function for determining if a Piece is in a Square.
func (p Piece) InSquare(s Square) bool {
	for _, pieceSquare := range p.Coords {
		if pieceSquare == s {
			return true
		}
	}
	return false
}

// InList is a function for determining if a Piece is in a slice of Squares.
// (this is O(n**2), so let's try not to use it
func (p Piece) InList(list []Square) bool {
	for _, pieceSquare := range p.Coords {
		for _, square := range list {
			if pieceSquare == square {
				return true
			}
		}
	}
	return false
}

// InPiece is a function for determining if a Piece is in a slice of Squares.
// (this is O(n**2), so let's try not to use it)
func (p Piece) InPiece(compare Piece) bool {
	for _, pieceSquare := range p.Coords {
		for _, square := range compare.Coords {
			if pieceSquare == square {
				return true
			}
		}
	}
	return false
}
