package hunter

import (
	"testing"

	"github.com/eaglerock1337/gobat/pkg/board"
)

func TestNewHunter(t *testing.T) {
	testHunter := NewHunter()

	// validate simple variables

	if testHunter.Turns != 0 {
		t.Errorf("NewHunter did not specify zero Turns taken")
	}

	if len(testHunter.Ships) != 5 {
		t.Errorf("NewHunter did not return all 5 Ship types")
	}

	if !testHunter.SeekMode {
		t.Errorf("NewHunter did not start in Seek Mode")
	}

	// validate piece data

	expectedData := map[int]int{2: 180, 3: 160, 4: 140, 5: 120}

	for ship, length := range expectedData {
		if testHunter.Data[ship].Len() != length {
			t.Errorf("PieceData for ship length %v did not return %v as expected, but %v", ship, length, testHunter.Data[ship].Len())
		}
	}

	for _, square := range testHunter.Shots {
		if square.Letter == 0 && square.Number == 0 {
			t.Errorf("NewHunter did not return a list of shots")
		}
	}

	// verify heatmap populated as expected

	expectedHeatMap := map[board.Square]int{
		{Letter: 0, Number: 0}: 10,
		{Letter: 2, Number: 2}: 28,
		{Letter: 4, Number: 4}: 34,
		{Letter: 0, Number: 2}: 19,
		{Letter: 0, Number: 4}: 22,
	}

	for square, heat := range expectedHeatMap {
		heatVal := testHunter.HeatMap[square.Letter][square.Number]
		if heatVal != heat {
			t.Errorf("HeatMap return unexpected value %v for square %v, got %v", heatVal, square.PrintSquare(), heat)
		}
	}
}

func TestDeleteShip(t *testing.T) {
	testDelete := NewHunter()
	result := testDelete.DeleteShip("Battleship")

	if result != nil {
		t.Errorf("DeleteShip returned an unexpected error: %v", result)
	}

	if len(testDelete.Ships) > 4 {
		t.Errorf("DeleteShip did not remove a ship, 5 ships still found")
	}
}

func TestBadDeleteShip(t *testing.T) {
	testBadDelete := NewHunter()
	firstResult := testBadDelete.DeleteShip("Cruiser")

	if firstResult != nil {
		t.Errorf("BadDeleteShip test not working: first DeleteShip returned unexpected error: %v", firstResult)
	}

	secondResult := testBadDelete.DeleteShip("Cruiser")

	if secondResult.Error() != "unable to find Ship to delete" {
		t.Errorf("DeleteShip did not error as expected on Ship array: %v", testBadDelete.Ships)
	}
}

func TestGetValidLengths(t *testing.T) {
	expectedSearchResults := []int{5, 4, 3, 2}
	testLengths := NewHunter()
	result := testLengths.GetValidLengths()

	if len(result) != len(expectedSearchResults) {
		t.Errorf("GetValidLengths was expected to return a slice of length %v, got %v", len(expectedSearchResults), len(result))
	}

	for i, v := range result {
		if v != expectedSearchResults[i] {
			t.Errorf("GetValidLengths did not return expected array %v, got %v", expectedSearchResults, result)
		}
	}
}

var exampleSquares = [5]board.Square{
	{Letter: 0, Number: 3},
	{Letter: 3, Number: 3},
	{Letter: 4, Number: 9},
	{Letter: 0, Number: 8},
	{Letter: 9, Number: 1},
}

func TestHitStack(t *testing.T) {
	testHitStack := NewHunter()

	for _, square := range exampleSquares {
		testHitStack.AddHitStack(square)
	}

	for i, val := range testHitStack.HitStack {
		if val != exampleSquares[i] {
			t.Errorf("HitStack did not return expected HitStack %v, got %v", expectedSquares, testHitStack.HitStack)
		}
	}
}

func TestDelHitStack(t *testing.T) {
	testDelHitStack := NewHunter()

	for _, square := range exampleSquares {
		testDelHitStack.AddHitStack(square)
	}

	deleteSquare := board.Square{Letter: 4, Number: 9}

	testDelHitStack.DelHitStack(deleteSquare)

	if len(testDelHitStack.HitStack) != len(exampleSquares)-1 {
		t.Errorf("DelHitStack did not delete %v as expected, got %v", deleteSquare, testDelHitStack.HitStack)
	}
}

var exampleWrongSquares = [5]board.Square{
	{Letter: 1, Number: 5},
	{Letter: 2, Number: 3},
	{Letter: 7, Number: 5},
	{Letter: 2, Number: 8},
	{Letter: 9, Number: 0},
}

func TestInHitStack(t *testing.T) {
	testInHitStack := NewHunter()

	for _, square := range exampleSquares {
		testInHitStack.AddHitStack(square)
	}

	for _, square := range exampleSquares {
		if !testInHitStack.InHitStack(square) {
			t.Errorf("InHitStack returned false for Square %v and HitStack %v, expected true", square, testInHitStack.HitStack)
		}
	}

	for _, square := range exampleWrongSquares {
		if testInHitStack.InHitStack(square) {
			t.Errorf("InHitStack returned true for Square %v and HitStack %v, expected false", square, testInHitStack.HitStack)
		}
	}
}

func TestAddShot(t *testing.T) {
	testAddShot := NewHunter()

	for _, square := range exampleSquares {
		testAddShot.AddShot(square)
	}

	for _, square := range exampleSquares {
		found := false
		for _, stackSquare := range testAddShot.Shots {
			if square == stackSquare {
				found = true
			}
		}
		if !found {
			t.Errorf("AddShot did not add %v to Shots array: %v", square, testAddShot.Shots)
		}
	}
}

func TestClearShots(t *testing.T) {
	testClearShots := NewHunter()

	for _, square := range exampleSquares {
		testClearShots.AddShot(square)
	}

	testClearShots.ClearShots()

	if len(testClearShots.Shots) > 0 {
		t.Errorf("ClearShots did not clear the Shots array, found: %v", testClearShots.Shots)
	}
}

var searchShips = [5]board.Ship{
	board.Ship("Destroyer"),
	board.Ship("Cruiser"),
	board.Ship("Submarine"),
	board.Ship("Battleship"),
	board.Ship("Carrier"),
}

var searchPieceSquares = [5][]board.Square{
	{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}},
	{{Letter: 2, Number: 3}, {Letter: 2, Number: 4}, {Letter: 2, Number: 5}},
	{{Letter: 6, Number: 1}, {Letter: 5, Number: 1}, {Letter: 4, Number: 1}},
	{{Letter: 9, Number: 3}, {Letter: 9, Number: 4}, {Letter: 9, Number: 5}, {Letter: 9, Number: 6}},
	{{Letter: 7, Number: 3}, {Letter: 6, Number: 3}, {Letter: 5, Number: 3}, {Letter: 4, Number: 3}, {Letter: 3, Number: 3}},
}

var expectedSearchResults = [5]board.Piece{
	{Type: searchShips[0], Coords: []board.Square{{Letter: 0, Number: 0}, {Letter: 0, Number: 1}}},
	{Type: searchShips[1], Coords: []board.Square{{Letter: 2, Number: 3}, {Letter: 2, Number: 4}, {Letter: 2, Number: 5}}},
	{Type: searchShips[2], Coords: []board.Square{{Letter: 4, Number: 1}, {Letter: 5, Number: 1}, {Letter: 6, Number: 1}}},
	{Type: searchShips[3], Coords: []board.Square{{Letter: 9, Number: 3}, {Letter: 9, Number: 4}, {Letter: 9, Number: 5}, {Letter: 9, Number: 6}}},
	{Type: searchShips[4], Coords: []board.Square{{Letter: 3, Number: 3}, {Letter: 4, Number: 3}, {Letter: 5, Number: 3}, {Letter: 6, Number: 3}, {Letter: 7, Number: 3}}},
}

func TestSearchPiece(t *testing.T) {
	for test, ship := range searchShips {
		testSearchPiece := NewHunter()
		numSquares := len(searchPieceSquares[test])

		for _, square := range searchPieceSquares[test] {
			testSearchPiece.AddShot(square)
			testSearchPiece.AddHitStack(square)
		}

		result, err := testSearchPiece.SearchPiece(searchPieceSquares[test][numSquares-1], ship)
		if err != nil {
			t.Errorf("SearchPiece failed and returned an error: %v", err)
		}

		if result.Type.GetType() != expectedSearchResults[test].Type.GetType() {
			t.Errorf("SearchPiece did not return ship type %v as expected, got %v", result.Type, expectedSearchResults[test].Type)
		}

		if result.Coords[0] != expectedSearchResults[test].Coords[0] || result.Coords[1] != expectedSearchResults[test].Coords[1] {
			t.Errorf("SearchPiece did not return result %v as expected, got %v", expectedSearchResults[test], result)
		}
	}
}

var badSearchSquares = [5]board.Square{
	{Letter: 6, Number: 4},
	{Letter: 3, Number: 3},
	{Letter: 7, Number: 1},
	{Letter: 9, Number: 2},
	{Letter: 3, Number: 8},
}

func TestBadSearchPiece(t *testing.T) {
	for test := range badSearchSquares {
		testBadSearchPiece := NewHunter()
		for _, square := range searchPieceSquares[test] {
			testBadSearchPiece.AddShot(square)
			testBadSearchPiece.AddHitStack(square)
		}

		result, err := testBadSearchPiece.SearchPiece(badSearchSquares[test], searchShips[test])

		if err == nil {
			t.Errorf("Error expected for search of %v, returned %v instead", badSearchSquares[test], result)
		}
	}
}

var dupeSearchPieceSquares = [5][]board.Square{
	{{Letter: 0, Number: 2}, {Letter: 0, Number: 0}, {Letter: 0, Number: 1}},
	{{Letter: 2, Number: 3}, {Letter: 2, Number: 4}, {Letter: 2, Number: 6}, {Letter: 2, Number: 7}, {Letter: 2, Number: 5}},
	{{Letter: 6, Number: 1}, {Letter: 5, Number: 1}, {Letter: 2, Number: 1}, {Letter: 3, Number: 1}, {Letter: 4, Number: 1}},
	{{Letter: 9, Number: 3}, {Letter: 9, Number: 4}, {Letter: 9, Number: 5}, {Letter: 9, Number: 7}, {Letter: 9, Number: 8}, {Letter: 9, Number: 9}, {Letter: 9, Number: 6}},
	{{Letter: 8, Number: 3}, {Letter: 7, Number: 3}, {Letter: 6, Number: 3}, {Letter: 5, Number: 3}, {Letter: 3, Number: 3}, {Letter: 2, Number: 3}, {Letter: 1, Number: 3}, {Letter: 0, Number: 3}, {Letter: 4, Number: 3}},
}

func TestDupeSearchPiece(t *testing.T) {
	for test, ship := range searchShips {
		testSearchPiece := NewHunter()
		numSquares := len(dupeSearchPieceSquares[test])

		for _, square := range dupeSearchPieceSquares[test] {
			testSearchPiece.AddShot(square)
			testSearchPiece.AddHitStack(square)
		}

		result, err := testSearchPiece.SearchPiece(dupeSearchPieceSquares[test][numSquares-1], ship)
		if err == nil {
			t.Errorf("TestDupeSearchPiece failed to error out: %v", result)
		}
	}
	// testSearchPiece.Dat
}
