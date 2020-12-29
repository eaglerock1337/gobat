package hunter

import (
	"testing"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

var testData = [10][10]int{
	{4, 2, 5, 23, 18, 90, 2, 0, 14, 3},
	{3, 48, 29, 2, 0, 23, 4, 8, 3, 12},
	{45, 23, 2, 0, 0, 0, 43, 23, 1, 4},
	{3, 45, 29, 2, 0, 23, 4, 8, 3, 12},
	{3, 48, 45, 2, 0, 2, 23, 8, 3, 45},
	{3, 2, 23, 2, 0, 45, 4, 47, 3, 19},
	{23, 48, 29, 2, 0, 7, 4, 8, 3, 18},
	{3, 29, 8, 2, 23, 6, 4, 18, 3, 17},
	{8, 45, 29, 3, 88, 8, 4, 8, 3, 29},
	{3, 45, 8, 2, 0, 23, 4, 8, 14, 12},
}

func TestInitialize(t *testing.T) {
	testHeatMap := HeatMap(testData)
	testHeatMap.Initialize()
	for i := range testHeatMap {
		for j := range testHeatMap[i] {
			if testHeatMap[i][j] != 0 {
				t.Errorf("Initialize did not zero the HeatMap at (%v, %v), got: %v", i, j, testHeatMap[i][j])
			}
		}
	}
}

var testSquares = [5]board.Square{
	{Letter: 2, Number: 5},
	{Letter: 8, Number: 3},
	{Letter: 4, Number: 5},
	{Letter: 4, Number: 5},
	{Letter: 5, Number: 4},
}

func TestAddSquare(t *testing.T) {
	var expected = map[board.Square]int{
		{Letter: 2, Number: 5}: 1,
		{Letter: 8, Number: 3}: 1,
		{Letter: 4, Number: 5}: 2,
		{Letter: 5, Number: 4}: 1,
		{Letter: 6, Number: 3}: 0,
	}
	var testmap HeatMap

	for _, value := range testSquares {
		testmap.AddSquare(value)
	}

	for square, amount := range expected {
		if testmap[square.Letter][square.Number] != amount {
			t.Errorf("AddSquare was incorrect for square %v, got: %v, want: %v", square, testmap[square.Letter][square.Number], amount)
		}
	}
}
