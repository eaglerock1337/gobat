package board

import "errors"

var values = map[string]int{
	"Empty": 0, "Miss": 1, "Destroyer": 2, "Submarine": 3,
	"Cruiser": 4, "Battleship": 5, "Carrier": 6, "Hit": 7,
}

var status = [8]string{
	"Empty", "Miss", "Destroyer", "Submarine",
	"Cruiser", "Battleship", "Carrier", "Hit",
}

// Board is a type for holding a standard 10x10 Battleship game board
type Board [10][10]int

// SetString sets a board value to a given string value
func (b *Board) SetString(s Square, value string) error {
	if val, ok := values[value]; ok {
		b[s.Letter][s.Number] = val
		return nil
	}
	return errors.New("Given value is not a valid value")
}

// SetInt sets a board value to a given integer value
func (b *Board) SetInt(s Square, value int) error {
	if value > 0 && value < 8 {
		b[s.Letter][s.Number] = value
		return nil
	}
	return errors.New("Given value out of range")
}

// GetString returns a given Square's string value
func (b Board) GetString(s Square) string {
	return status[b[s.Letter][s.Number]]
}

// GetInt returns a given Square's integer value
func (b Board) GetInt(s Square) int {
	return b[s.Letter][s.Number]
}

// IsEmpty returns whether a given Square is empty
func (b Board) IsEmpty(s Square) bool {
	return (b[s.Letter][s.Number] == 0)
}

// IsMiss returns whether a given Square is a miss
func (b Board) IsMiss(s Square) bool {
	return (b[s.Letter][s.Number] == 1)
}

// IsHit returns whether a given square is any type of hit
func (b Board) IsHit(s Square) bool {
	return (b[s.Letter][s.Number] > 1)
}

// IsUnsunk returns whether a given square is to an unsunk ship
func (b Board) IsUnsunk(s Square) bool {
	return (b[s.Letter][s.Number] == 7)
}

// IsSunk returns whether a given square is to any sunk ship
func (b Board) IsSunk(s Square) bool {
	val := b[s.Letter][s.Number]
	return (val > 1 && val < 7)
}

// IsPiece returns whether a square belongs to a specific piece
// func (b Board) IsShip(s Square, p Piece) bool {
// 	return (status[b[s.Letter][s.Number]] == p.String)
// }
