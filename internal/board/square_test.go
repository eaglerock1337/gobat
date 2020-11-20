package board

import "testing"

func TestLetter(t *testing.T) {
	example := Square{2, 2}
	result := "C"

	answer := example.Letter()
	if answer != result {
		t.Errorf("Letter function was incorrect, got: %v, want:%v", answer, result)
	}
}

func TestNumber(t *testing.T) {
	example := Square{2, 2}
	result := "3"

	answer := example.Number()
	if answer != result {
		t.Errorf("Number function was incorrect, got: %v, want:%v", answer, result)
	}
}

func TestSquare(t *testing.T) {
	example := Square{2, 2}
	result := "C3"

	answer := example.Square()
	if answer != result {
		t.Errorf("Square function was incorrect, got: %v, want:%v", answer, result)
	}
}
