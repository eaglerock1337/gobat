package board

import "testing"

var examples = []Square{
	Square{0, 0},
	Square{2, 2},
	Square{5, 2},
	Square{9, 6},
	Square{6, 9},
}

func TestLetter(t *testing.T) {
	results := [5]string{"A", "C", "F", "J", "G"}

	for i := 0; i < 5; i++ {
		answer := examples[i].Letter()
		if answer != results[i] {
			t.Errorf("Letter function was incorrect, got: %v, want:%v", answer, results[i])
		}
	}
}

func TestNumber(t *testing.T) {
	results := [5]string{"1", "3", "3", "7", "10"}

	for i := 0; i < 5; i++ {
		answer := examples[i].Number()
		if answer != results[i] {
			t.Errorf("Letter function was incorrect, got: %v, want:%v", answer, results[i])
		}
	}
}

func TestSquare(t *testing.T) {
	results := [5]string{"A1", "C3", "F3", "J7", "G10"}

	for i := 0; i < 4; i++ {
		answer := examples[i].Square()
		if answer != results[i] {
			t.Errorf("Square function was incorrect, got: %v, want:%v", answer, results[i])
		}
	}
}
