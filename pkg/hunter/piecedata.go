package hunter

import (
	"errors"

	"github.com/eaglerock1337/gobat/pkg/board"
)

// PieceData is a slice of pieces that contains all possible placements for
// a piece with a given length.
type PieceData []board.Piece

// GenPieceData generates a complete heatdata for a given Ship.
// I should probably add error checking into this.
func GenPieceData(ship board.Ship) PieceData {
	var data PieceData
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10-ship.GetLength(); j++ {
			// Add the ship horizontally, errors can be ignored due to constraints
			hSquare, _ := board.SquareByValue(j, i)
			hPiece, _ := board.NewPiece(ship, hSquare, true)
			data = append(data, hPiece)

			// Add the ship vertically, errors can be ignored due to constraints
			vSquare, _ := board.SquareByValue(i, j)
			vPiece, _ := board.NewPiece(ship, vSquare, false)
			data = append(data, vPiece)
		}
	}
	return data
}

// Remove removes a piece in the piece data slice.
func (p *PieceData) Remove(pos int) error {
	lastPos := len(*p) - 1
	if pos < 0 || pos > lastPos {
		return errors.New("position to remove is out of bounds")
	}

	(*p)[pos] = (*p)[lastPos]
	*p = (*p)[:lastPos]
	return nil
}

// DeleteSquare removes all Pieces that reside in a given Square.
func (p *PieceData) DeleteSquare(square board.Square) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].InSquare(square) {
			p.Remove(i)
			i-- // reiterate over i since it was overwritten by DeletePiece()
		}
	}
}

// DeletePiece removes all references to all squares in a given Piece's position.
func (p *PieceData) DeletePiece(piece board.Piece) {
	for _, square := range piece.Coords {
		p.DeleteSquare(square)
	}
}

// Len returns the length of the PieceData slice.
func (p *PieceData) Len() int {
	return len(*p)
}
