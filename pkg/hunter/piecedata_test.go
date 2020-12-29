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

var expectedSquares = [5][5][]board.Square{
	{
		{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}},
		{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}, {Letter: 0, Number: 4}},
		{{Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}, {Letter: 5, Number: 0}},
		{{Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}, {Letter: 0, Number: 4}, {Letter: 0, Number: 5}},
		{{Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}, {Letter: 5, Number: 0}, {Letter: 6, Number: 0}},
	},
	{
		{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}},
		{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}},
		{{Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}},
		{{Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}, {Letter: 0, Number: 4}},
		{{Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}, {Letter: 5, Number: 0}},
	},
	{
		{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}, {Letter: 2, Number: 0}},
		{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}, {Letter: 0, Number: 2}},
		{{Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}},
		{{Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}},
		{{Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}},
	},
	{
		{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}, {Letter: 2, Number: 0}},
		{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}, {Letter: 0, Number: 2}},
		{{Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}},
		{{Letter: 0, Number: 1}, {Letter: 0, Number: 2}, {Letter: 0, Number: 3}},
		{{Letter: 2, Number: 0}, {Letter: 3, Number: 0}, {Letter: 4, Number: 0}},
	},
	{
		{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}},
		{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}},
		{{Letter: 1, Number: 0}, {Letter: 2, Number: 0}},
		{{Letter: 0, Number: 1}, {Letter: 0, Number: 2}},
		{{Letter: 2, Number: 0}, {Letter: 3, Number: 0}},
	},
}

var expectedValues = [5][5]board.Piece{
	{
		{Type: exampleShips[0], Coords: expectedSquares[0][0]},
		{Type: exampleShips[0], Coords: expectedSquares[0][1]},
		{Type: exampleShips[0], Coords: expectedSquares[0][2]},
		{Type: exampleShips[0], Coords: expectedSquares[0][3]},
		{Type: exampleShips[0], Coords: expectedSquares[0][4]},
	},
	{
		{Type: exampleShips[1], Coords: expectedSquares[1][0]},
		{Type: exampleShips[1], Coords: expectedSquares[1][1]},
		{Type: exampleShips[1], Coords: expectedSquares[1][2]},
		{Type: exampleShips[1], Coords: expectedSquares[1][3]},
		{Type: exampleShips[1], Coords: expectedSquares[1][4]},
	},
	{
		{Type: exampleShips[2], Coords: expectedSquares[2][0]},
		{Type: exampleShips[2], Coords: expectedSquares[2][1]},
		{Type: exampleShips[2], Coords: expectedSquares[2][2]},
		{Type: exampleShips[2], Coords: expectedSquares[2][3]},
		{Type: exampleShips[2], Coords: expectedSquares[2][4]},
	},
	{
		{Type: exampleShips[3], Coords: expectedSquares[3][0]},
		{Type: exampleShips[3], Coords: expectedSquares[3][1]},
		{Type: exampleShips[3], Coords: expectedSquares[3][2]},
		{Type: exampleShips[3], Coords: expectedSquares[3][3]},
		{Type: exampleShips[3], Coords: expectedSquares[3][4]},
	},
	{
		{Type: exampleShips[4], Coords: expectedSquares[4][0]},
		{Type: exampleShips[4], Coords: expectedSquares[4][1]},
		{Type: exampleShips[4], Coords: expectedSquares[4][2]},
		{Type: exampleShips[4], Coords: expectedSquares[4][3]},
		{Type: exampleShips[4], Coords: expectedSquares[4][4]},
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

var exampleRemoveData = [10][]board.Square{
	{{Letter: 0, Number: 0}, {Letter: 1, Number: 0}, {Letter: 2, Number: 0}, {Letter: 3, Number: 0}},
	{{Letter: 3, Number: 0}, {Letter: 4, Number: 0}, {Letter: 5, Number: 0}, {Letter: 6, Number: 0}},
	{{Letter: 3, Number: 2}, {Letter: 4, Number: 2}, {Letter: 5, Number: 2}, {Letter: 6, Number: 2}},
	{{Letter: 6, Number: 0}, {Letter: 7, Number: 0}, {Letter: 8, Number: 0}, {Letter: 9, Number: 0}},
	{{Letter: 6, Number: 2}, {Letter: 7, Number: 2}, {Letter: 8, Number: 2}, {Letter: 9, Number: 2}},
	{{Letter: 4, Number: 5}, {Letter: 5, Number: 5}, {Letter: 6, Number: 5}, {Letter: 7, Number: 5}},
	{{Letter: 7, Number: 3}, {Letter: 7, Number: 4}, {Letter: 7, Number: 5}, {Letter: 7, Number: 6}},
	{{Letter: 9, Number: 5}, {Letter: 9, Number: 6}, {Letter: 9, Number: 7}, {Letter: 9, Number: 8}},
	{{Letter: 5, Number: 1}, {Letter: 5, Number: 2}, {Letter: 5, Number: 3}, {Letter: 5, Number: 4}},
	{{Letter: 5, Number: 3}, {Letter: 5, Number: 4}, {Letter: 5, Number: 5}, {Letter: 5, Number: 6}},
}

func TestRemove(t *testing.T) {
	var exampleData = PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[0]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[1]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[2]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[4]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[5]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[7]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[9]},
	}
	placesToRemove := [5]int{2, 4, 5, 1, 2}
	answer := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[0]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[7]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
	}

	for _, value := range placesToRemove {
		err := exampleData.Remove(value)
		if err != nil {
			t.Errorf("Remove(%v) returned an error trying to run on PieceData of length %v: %v", value, len(exampleData), err)
		}
	}

	for i, piece := range exampleData {
		for j, square := range piece.Coords {
			if square != answer[i].Coords[j] {
				t.Errorf("Remove function was incorrect, got: %v, want: %v", piece.Coords, answer[i].Coords)
				break
			}
		}
	}
}

func TestBadRemove(t *testing.T) {
	placesToRemove := [5]int{6, 8, 5, -1, 44}
	samplePieceData := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[0]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[7]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
	}

	for _, value := range placesToRemove {
		err := samplePieceData.Remove(value)
		if err == nil {
			t.Errorf("Remove did not error as expected with %v, returned PieceData: %v", value, samplePieceData)
		}
	}
}

func TestDeleteSquare(t *testing.T) {
	squaresToRemove := [5]board.Square{
		{Letter: 3, Number: 0},
		{Letter: 7, Number: 4},
		{Letter: 6, Number: 8},
		{Letter: 9, Number: 5},
		{Letter: 5, Number: 3},
	}
	exampleData := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[0]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[1]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[2]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[4]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[5]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[7]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[9]},
	}
	expected := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[5]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[4]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[2]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
	}

	for _, square := range squaresToRemove {
		exampleData.DeleteSquare(square)
	}

	for i, piece := range exampleData {
		for j, square := range piece.Coords {
			if square != expected[i].Coords[j] {
				t.Errorf("DeleteSquare function was incorrect, got: %v, want: %v", piece.Coords, expected[i].Coords)
				break
			}
		}
	}
}

func TestDeletePiece(t *testing.T) {
	piecesToRemove := [5]board.Piece{
		{Type: board.Ship("Cruiser"), Coords: []board.Square{{Letter: 3, Number: 0}, {Letter: 3, Number: 1}}},
		{Type: board.Ship("Cruiser"), Coords: []board.Square{{Letter: 4, Number: 4}, {Letter: 4, Number: 5}}},
		{Type: board.Ship("Cruiser"), Coords: []board.Square{{Letter: 6, Number: 0}, {Letter: 6, Number: 1}}},
		{Type: board.Ship("Cruiser"), Coords: []board.Square{{Letter: 7, Number: 2}, {Letter: 8, Number: 2}}},
		{Type: board.Ship("Cruiser"), Coords: []board.Square{{Letter: 8, Number: 5}, {Letter: 9, Number: 5}}},
	}
	exampleData := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[0]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[1]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[2]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[3]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[4]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[5]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[7]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[9]},
	}
	expected := PieceData{
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[9]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[8]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[2]},
		{Type: board.Ship("Battleship"), Coords: exampleRemoveData[6]},
	}

	for _, piece := range piecesToRemove {
		exampleData.DeletePiece(piece)
	}

	for i, piece := range exampleData {
		for j, square := range piece.Coords {
			if square != expected[i].Coords[j] {
				t.Errorf("DeletePiece function was incorrect, got: %v, want: %v", piece.Coords, expected[i].Coords)
				break
			}
		}
	}
}
