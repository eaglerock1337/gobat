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

	if secondResult.Error() != "Ship not found" {
		t.Errorf("BadDeleteShip did not error as expected on Ship array: %v", testBadDelete.Ships)
	}
}

func TestGetValidLengths(t *testing.T) {
	expectedResult := []int{5, 4, 3, 2}
	testLengths := NewHunter()
	result := testLengths.GetValidLengths()

	if len(result) != len(expectedResult) {
		t.Errorf("GetValidLengths was expected to return a slice of length %v, got %v", len(expectedResult), len(result))
	}

	for i, v := range result {
		if v != expectedResult[i] {
			t.Errorf("GetValidLengths did not return expected array %v, got %v", expectedResult, result)
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
			t.Errorf("TestHitStack did not return expected HitStack %v, got %v", expectedSquares, testHitStack.HitStack)
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
		t.Errorf("TestDelHitStack did not delete %v as expected, got %v", deleteSquare, testDelHitStack.HitStack)
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
			t.Errorf("TestAddShot did not find %v in Shots array: %v", square, testAddShot.Shots)
		}
	}
}

// func TestSearchPiece(t *testing.T) {
// 	testSearchPiece := NewHunter()
// 	testSearchPiece.AddShot(board.Square{Letter: 0, Number: 0})
// 	testSearchPiece.AddShot(board.Square{Letter: 0, Number: 1})
// 	testSearchPiece.SearchPiece
// }
