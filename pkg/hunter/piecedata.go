package hunter

import (
	"github.com/eaglerock1337/go/battleship/pkg/board"
)

// PieceData is a slice of pieces that contains all possible placements for
// a piece with a given length.
type PieceData []board.Piece

// GenPieceData generates a complete heatdata for a given Ship.
func GenPieceData(ship board.Ship) PieceData {
	var data PieceData
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10-ship.GetLength(); j++ {
			// Add the ship horizontally, errors can be ignored due to constraints
			hSquare, _ := board.SquareByValue(i, j)
			hPiece, _ := board.NewPiece(ship, hSquare, true)
			data = append(data, hPiece)

			// Add the ship vertically, errors can be ignored due to constraints
			vSquare, _ := board.SquareByValue(j, i)
			vPiece, _ := board.NewPiece(ship, vSquare, false)
			data = append(data, vPiece)
		}
	}
	return data
}

// DeletePiece removes a piece in the piece data slice.
func (p *PieceData) DeletePiece(pos int) {
	lastPos := len(*p) - 1
	(*p)[pos] = (*p)[lastPos]
	*p = (*p)[:lastPos-1]
}

// DeleteSquare removes all Pieces that reside in a given Square.
func (p *PieceData) DeleteSquare(square board.Square) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].InSquare(square) {
			p.DeletePiece(i)
			i-- // reiterate over i since it was overwritten by DeletePiece()
		}
	}
}
