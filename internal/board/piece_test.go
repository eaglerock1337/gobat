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
	{0, 0},
	{2, 2},
	{5, 2},
	{9, 6},
	{6, 9},
}

var examplePieces = [5]Piece{
	{Ship("Carrier"), []Square{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}},
	{Ship("Battleship"), []Square{{2, 2}, {2, 3}, {2, 4}, {2, 5}}},
	{Ship("Cruiser"), []Square{{5, 2}, {6, 2}, {7, 2}}},
	{Ship("Submarine"), []Square{{9, 6}, {9, 7}, {9, 8}}},
	{Ship("Destroyer"), []Square{{6, 9}, {7, 9}}},
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

func TestShipTypes(t *testing.T) {
	answer := ShipTypes()
	for i, input := range answer {
		if input != exampleShips[i] {
			t.Errorf("ShipType was incorrect, got: %v, want: %v", input, exampleShips[i])
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
