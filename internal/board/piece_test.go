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

func TestNewShip(t *testing.T) {
	for i, input := range exampleTypes {
		answer, err := NewShip(input)

		if err != nil {
			t.Errorf("NewShip returned an error: %v", err)
		} else if answer != exampleShips[i] {
			t.Errorf("NewShip function was incorrect, got: %v, want:%v", answer, exampleShips[i])
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
