/*
Package board implements structs and custom types that can be used for representing
and handling a Battleship board and its pieces. All custom types have built-in
error handling and member functions for referencing the data in multiple ways.

The structs and types are as follows:

board.Square - a struct that contains a single coordinate on a Battleship board
board.Ship   - a custom string that represents a type of ship in Battleship
board.Piece  - a struct that contains a ship's type and coordinates on the board
board.Board  - a 10x10 Battleship board with standard letter-number referencing

With the exception of board.Board (which returns a zeroed array of integers),
each type/struct has one or more functions for creation with validation and with
error reporting to make input validation easy. All member functions that update
data will also report error statuses for the same reason.
*/
package board

import "errors"

// Laziness wins the day
var values = map[string]int{
	"Empty": 0, "Miss": 1, "Destroyer": 2, "Submarine": 3,
	"Cruiser": 4, "Battleship": 5, "Carrier": 6, "Hit": 7,
}

// I can't tell if this is genius or really stupid
var status = [8]string{
	"Empty", "Miss", "Destroyer", "Submarine",
	"Cruiser", "Battleship", "Carrier", "Hit",
}

// Board is a type for holding a standard 10x10 Battleship game board
type Board [10][10]int

// Board update methods

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

// Board retrieval methods

// GetString returns a given Square's string value
func (b Board) GetString(s Square) string {
	return status[b[s.Letter][s.Number]]
}

// GetInt returns a given Square's integer value
func (b Board) GetInt(s Square) int {
	return b[s.Letter][s.Number]
}

// Board boolean methods

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

// IsShip returns whether a square belongs to a specific ship type
func (b Board) IsShip(s Square, sh Ship) bool {
	return (status[b[s.Letter][s.Number]] == string(sh))
}
