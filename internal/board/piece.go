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
	Type   string
	Coords []Square
}

// PieceByVector defines a Piece by the following variables:
// shipType: a string with the desired ship type
// startSquare: the starting Square of the ship from the topmost, leftmost square
// horizontal: a boolean that determines either horizontal or vertical placement
func PieceByVector(shipType string, startSquare Square, horizontal bool) (Piece, error) {
	if squares, ok := ships[shipType]; ok {
		// TODO: do piece on board validation here
		var newPiece Piece
		newPiece.Type = shipType
		newPiece.Coords = make([]Square, squares)
		return newPiece, nil
	}
	return Piece{}, errors.New("Something went wrong")
}
