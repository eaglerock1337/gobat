package hunter

import (
	"testing"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

var exampleShips = [5]board.Ship{
	board.Ship("Carrier"),
	board.Ship("Battleship"),
	board.Ship("Cruiser"),
	board.Ship("Submarine"),
	board.Ship("Destroyer"),
}

var expectedLengths = [5]int{120, 140, 160, 160, 180}

var expectedValues = [5][5]board.Piece{
	{
		{board.Ship("Carrier"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}},
		{board.Ship("Carrier"), []board.Square{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}},
		{board.Ship("Carrier"), []board.Square{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}}},
		{board.Ship("Carrier"), []board.Square{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}},
		{board.Ship("Carrier"), []board.Square{{2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}}},
	},
	{
		{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		{board.Ship("Battleship"), []board.Square{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		{board.Ship("Battleship"), []board.Square{{1, 0}, {2, 0}, {3, 0}, {4, 0}}},
		{board.Ship("Battleship"), []board.Square{{0, 1}, {0, 2}, {0, 3}, {0, 4}}},
		{board.Ship("Battleship"), []board.Square{{2, 0}, {3, 0}, {4, 0}, {5, 0}}},
	},
	{
		{board.Ship("Cruiser"), []board.Square{{0, 0}, {1, 0}, {2, 0}}},
		{board.Ship("Cruiser"), []board.Square{{0, 0}, {0, 1}, {0, 2}}},
		{board.Ship("Cruiser"), []board.Square{{1, 0}, {2, 0}, {3, 0}}},
		{board.Ship("Cruiser"), []board.Square{{0, 1}, {0, 2}, {0, 3}}},
		{board.Ship("Cruiser"), []board.Square{{2, 0}, {3, 0}, {4, 0}}},
	},
	{
		{board.Ship("Submarine"), []board.Square{{0, 0}, {1, 0}, {2, 0}}},
		{board.Ship("Submarine"), []board.Square{{0, 0}, {0, 1}, {0, 2}}},
		{board.Ship("Submarine"), []board.Square{{1, 0}, {2, 0}, {3, 0}}},
		{board.Ship("Submarine"), []board.Square{{0, 1}, {0, 2}, {0, 3}}},
		{board.Ship("Submarine"), []board.Square{{2, 0}, {3, 0}, {4, 0}}},
	},
	{
		{board.Ship("Destroyer"), []board.Square{{0, 0}, {1, 0}}},
		{board.Ship("Destroyer"), []board.Square{{0, 0}, {0, 1}}},
		{board.Ship("Destroyer"), []board.Square{{1, 0}, {2, 0}}},
		{board.Ship("Destroyer"), []board.Square{{0, 1}, {0, 2}}},
		{board.Ship("Destroyer"), []board.Square{{2, 0}, {3, 0}}},
	},
}

func TestGenPieceData(t *testing.T) {
	for i, input := range exampleShips {
		answer := GenPieceData(input)
		for j, piece := range expectedValues[i] {
			for k, square := range piece.Coords {
				if square != answer[j].Coords[k] {
					t.Errorf("GenPieceData function was incorrect, got: %v, want: %v", answer[j].Coords, piece.Coords)
					break
				}
			}

			if piece.Type != answer[j].Type {
				t.Errorf("GenPieceData function was incorrect, got: %v, want: %v", answer[j], piece)
			}
		}

		if expectedLengths[i] != len(answer) {
			t.Errorf("GenPieceData function was incorrect, got: %v length, want: %v length", len(answer), expectedLengths[i])
		}
	}
}

var examplePieceData = PieceData[]{
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{1, 0}, {2, 0}, {3, 0}, {4, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{1, 1}, {2, 1}, {3, 1}, {4, 1}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{6, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	board.Piece{board.Ship("Battleship"), []board.Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
}
