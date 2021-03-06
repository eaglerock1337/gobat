package board

import (
	"testing"
)

var boardSquares = [5]Square{
	{0, 0},
	{2, 2},
	{5, 2},
	{9, 6},
	{6, 9},
}

var boardStrings = [5]string{
	"Miss", "Destroyer", "Battleship", "Carrier", "Hit",
}

func TestSetString(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		err := testboard.SetString(input, boardStrings[i])
		answer := testboard[input.Letter][input.Number]

		if err != nil {
			t.Errorf("SetString returned an error: %v", err)
		} else if answer != boardIntegers[i] {
			t.Errorf("SetString function was incorrect, got: %v, want: %v", answer, boardIntegers[i])
		}
	}
}

var badBoardStrs = [5]string{
	"Mister", "destroyer", "Your mom", "Poop", "Mah boat",
}

func TestBadSetString(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		err := testboard.SetString(input, badBoardStrs[i])

		if err == nil {
			t.Errorf(
				"SetString did not error as expected with %v, set value: %v",
				badBoardStrs[i],
				testboard[input.Letter][input.Number],
			)
		}
	}
}

var boardIntegers = [5]int{1, 2, 5, 6, 7}

func TestSetInt(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		err := testboard.SetInt(input, boardIntegers[i])
		answer := testboard[input.Letter][input.Number]

		if err != nil {
			t.Errorf("SetInt returned an error: %v", err)
		} else if answer != boardIntegers[i] {
			t.Errorf("SetInt function was incorrect, got: %v, want: %v", answer, boardIntegers[i])
		}
	}
}

var badBoardInts = [5]int{-1, 8, 20, -3, 17}

func TestBadSetInt(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		err := testboard.SetInt(input, badBoardInts[i])

		if err == nil {
			t.Errorf(
				"SetInt did not error as expected with %v, set value: %v",
				badBoardInts[i],
				testboard[input.Letter][input.Number],
			)
		}
	}
}

var boardTestPieces = [5]Piece{
	{Ship("Carrier"), []Square{{0, 7}, {1, 7}, {2, 7}, {3, 7}, {4, 7}}},
	{Ship("Battleship"), []Square{{7, 2}, {7, 3}, {7, 4}, {7, 5}}},
	{Ship("Cruiser"), []Square{{5, 8}, {6, 8}, {7, 8}}},
	{Ship("Submarine"), []Square{{9, 6}, {9, 7}, {9, 8}}},
	{Ship("Destroyer"), []Square{{6, 9}, {7, 9}}},
}

var boardPieceValues = [5]int{6, 5, 4, 3, 2}

func TestSetPiece(t *testing.T) {
	var testboard Board

	for i, input := range boardTestPieces {
		testboard.SetPiece(input)
		testValue := testboard[input.Coords[0].Letter][input.Coords[0].Number]
		if testValue != boardPieceValues[i] {
			t.Errorf("SetPiece function was incorrect, got: %v, want: %v", testValue, boardPieceValues[i])
		}
	}
}

func TestPlacePiece(t *testing.T) {
	var testboard Board

	for i, input := range boardTestPieces {
		err := testboard.PlacePiece(input)
		testvalue := testboard[input.Coords[0].Letter][input.Coords[0].Number]

		if err != nil {
			t.Errorf("SetInt returned an error: %v", err)
		} else if testvalue != boardPieceValues[i] {
			t.Errorf("PlacePiece function was incorrect, got: %v, want: %v", testvalue, boardPieceValues[i])
		}
	}
}

var boardPieceSquares = [5]Square{
	{0, 7},
	{7, 2},
	{5, 8},
	{9, 6},
	{6, 9},
}

func TestBadPlacePiece(t *testing.T) {
	var testboard Board

	for _, square := range boardPieceSquares {
		testboard[square.Letter][square.Number] = 7
	}

	for _, input := range boardTestPieces {
		err := testboard.PlacePiece(input)
		if err == nil {
			t.Errorf(
				"PlacePiece did not error as expected with %v, set first square: %v",
				input,
				testboard[input.Coords[0].Letter][input.Coords[0].Number],
			)
		}
	}
}

func TestGetString(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardIntegers[i]
		answer := testboard.GetString(input)

		if answer != boardStrings[i] {
			t.Errorf("GetString function was incorrect, got: %v, want: %v", answer, boardStrings[i])
		}
	}
}

func TestGetInt(t *testing.T) {
	var testboard Board

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardIntegers[i]
		answer := testboard.GetInt(input)

		if answer != boardIntegers[i] {
			t.Errorf("GetInt function was incorrect, got: %v, want: %v", answer, boardIntegers[i])
		}
	}
}

var boardTestVals = [5]int{0, 1, 2, 4, 7}

func TestIsEmpty(t *testing.T) {
	var testboard Board
	var expected = [5]bool{true, false, false, false, false}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsEmpty(input)

		if answer != expected[i] {
			t.Errorf("IsEmpty function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}

func TestIsMiss(t *testing.T) {
	var testboard Board
	var expected = [5]bool{false, true, false, false, false}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsMiss(input)

		if answer != expected[i] {
			t.Errorf("IsMiss function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}

func TestIsHit(t *testing.T) {
	var testboard Board
	var expected = [5]bool{false, false, true, true, true}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsHit(input)

		if answer != expected[i] {
			t.Errorf("IsHit function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}

func TestIsUnsunk(t *testing.T) {
	var testboard Board
	var expected = [5]bool{false, false, false, false, true}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsUnsunk(input)

		if answer != expected[i] {
			t.Errorf("IsUnsunk function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}

func TestIsSunk(t *testing.T) {
	var testboard Board
	var expected = [5]bool{false, false, true, true, false}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsSunk(input)

		if answer != expected[i] {
			t.Errorf("IsSunk function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}

var boardTestShips = [5]Ship{
	Ship("Destroyer"),
	Ship("Destroyer"),
	Ship("Battleship"),
	Ship("Cruiser"),
	Ship("Submarine"),
}

func TestIsShip(t *testing.T) {
	var testboard Board
	var expected = [5]bool{false, false, false, true, false}

	for i, input := range boardSquares {
		testboard[input.Letter][input.Number] = boardTestVals[i]
		answer := testboard.IsShip(input, boardTestShips[i])

		if answer != expected[i] {
			t.Errorf("IsShip function was incorrect, got: %v, want: %v", answer, expected[i])
		}
	}
}
