package board

import (
	"testing"
)

var exampleTypes = [5]string{
	"Carrier",
	"Battleship",
	"Cruiser",
	"Submarine",
	"Destroyer",
}

var exampleShips = [5]Ship{
	Ship("Carrier"),
	Ship("Battleship"),
	Ship("Cruiser"),
	Ship("Submarine"),
	Ship("Destroyer"),
}

var exampleSizes = [5]int{5, 4, 3, 3, 2}

var badShips = [5]string{
	"Your Mom",
	"destroyer",
	"BATTLESHIP",
	"SuBmArInE",
	"I'm on a boat",
}

var pieceSquares = [5]Square{
	{0, 7},
	{7, 2},
	{5, 8},
	{9, 6},
	{6, 9},
}

var examplePieces = [5]Piece{
	{Ship("Carrier"), []Square{{0, 7}, {1, 7}, {2, 7}, {3, 7}, {4, 7}}},
	{Ship("Battleship"), []Square{{7, 2}, {7, 3}, {7, 4}, {7, 5}}},
	{Ship("Cruiser"), []Square{{5, 8}, {6, 8}, {7, 8}}},
	{Ship("Submarine"), []Square{{9, 6}, {9, 7}, {9, 8}}},
	{Ship("Destroyer"), []Square{{6, 9}, {7, 9}}},
}

var testSquares = [5]Square{{2, 7}, {3, 3}, {7, 8}, {9, 5}, {6, 9}}

var testLists = [5][]Square{
	{{0, 7}, {3, 3}, {4, 5}},
	{{3, 4}, {2, 9}, {8, 0}, {1, 3}, {4, 4}},
	{{4, 3}, {4, 4}, {4, 4}, {5, 1}, {6, 2}},
	{{0, 4}, {6, 6}, {7, 8}, {9, 5}, {9, 6}, {8, 9}},
	{{5, 9}, {6, 8}, {7, 8}, {8, 8}, {9, 8}, {7, 7}},
}

var testPieces = [5]Piece{
	{Ship("Destroyer"), []Square{{2, 7}, {2, 8}}},
	{Ship("Submarine"), []Square{{0, 6}, {1, 6}, {2, 6}}},
	{Ship("Cruiser"), []Square{{2, 2}, {3, 2}, {4, 2}}},
	{Ship("Destroyer"), []Square{{8, 7}, {9, 7}}},
	{Ship("Destroyer"), []Square{{5, 9}, {6, 9}}},
}

func TestNewShip(t *testing.T) {
	for i, input := range exampleTypes {
		answer, err := NewShip(input)

		if err != nil {
			t.Errorf("NewShip returned an error: %v", err)
		} else if answer != exampleShips[i] {
			t.Errorf("NewShip function was incorrect, got: %v, want: %v", answer, exampleShips[i])
		}
	}
}

func TestBadNewShip(t *testing.T) {
	for _, input := range badShips {
		answer, err := NewShip(input)

		if err == nil {
			t.Errorf("NewShip did not error as expected with %v, returned Ship: %v", input, answer)
		}
	}
}

func TestShipTypes(t *testing.T) {
	answer := ShipTypes()
	for i, input := range answer {
		if input != exampleShips[i] {
			t.Errorf("ShipType was incorrect, got: %v, want: %v", input, exampleShips[i])
		}
	}
}

func TestShipType(t *testing.T) {
	for i, input := range exampleShips {
		answer := input.Type()

		if answer != exampleTypes[i] {
			t.Errorf("Type was incorrect, got: %v, want: %v", answer, exampleStrings[i])
		}
	}
}

func TestShipLength(t *testing.T) {
	for i, input := range exampleShips {
		answer := input.Length()

		if answer != exampleSizes[i] {
			t.Errorf("Length was incorrect, got %v, want: %v", answer, exampleSizes[i])
		}
	}
}

func TestNewPiece(t *testing.T) {
	horizontal := [5]bool{true, false, true, false, true}

	for i, input := range exampleShips {
		answer, err := NewPiece(input, pieceSquares[i], horizontal[i])

		if err != nil {
			t.Errorf("NewPiece returned an error: %v", err)
		} else {
			if answer.Type != examplePieces[i].Type {
				t.Errorf("NewPiece set type incorrect, got: %v, want: %v", answer.Type, examplePieces[i].Type)
			}
			for j, coord := range answer.Coords {
				if coord != examplePieces[i].Coords[j] {
					t.Errorf("NewPiece coordinates were incorrect, got: %v, want: %v", coord, examplePieces[i].Coords[j])
				}
			}
		}
	}
}

func TestBadNewPiece(t *testing.T) {
	horizontal := [5]bool{false, true, false, true, false}

	for i, input := range exampleShips {
		answer, err := NewPiece(input, pieceSquares[i], horizontal[i])

		if err == nil {
			t.Errorf("NewPiece did not error as expected with %v, returned Piece: %v", input, answer)
		}
	}
}

func TestInSquare(t *testing.T) {
	expected := [5]bool{true, false, true, false, true}

	for i, input := range examplePieces {
		answer := input.InSquare(testSquares[i])

		if answer != expected[i] {
			t.Errorf("InSquare was incorrect with: %v, want: %v", input, expected[i])
		}
	}
}

func TestInList(t *testing.T) {
	expected := [5]bool{true, false, false, true, false}

	for i, input := range examplePieces {
		answer := input.InList(testLists[i])

		if answer != expected[i] {
			t.Errorf("InList was incorrect with: %v, want: %v", input, expected[i])
		}
	}
}

func TestInPiece(t *testing.T) {
	expected := [5]bool{true, false, false, true, true}

	for i, input := range examplePieces {
		answer := input.InPiece(testPieces[i])

		if answer != expected[i] {
			t.Errorf("InPiece was incorrect with: %v, want: %v", input, expected[i])
		}
	}
}
