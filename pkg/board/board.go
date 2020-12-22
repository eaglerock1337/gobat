/*
Package board implements structs and custom types that can be used for representing
and handling a Battleship board and its pieces. All custom types have built-in
error handling and member functions for referencing the data in multiple ways.

The Square struct and custom Ship string type are supporting types that allow for
easy input validation for referencing a board coordinate (e.g. A1, C3) or for
a type of ship (e.g. Cruiser, Battleship). These types should only be created
through their creation methids to provide validation feedback via errors.

The Piece struct utilizes the above two types to allow a ship's potential location
on the board to be defined in a way that provides automatic input validation. The
member methods of the above two types allow for Piece to be used with Board easily
by referencing coordinates by its zero-based array location. Just like the above
types, Piece should only be created by its creation method for validation purposes.

Board is a simple 10x10 array of integers that can be defined normally without a
creation function ("var myBoard Board"). The above types can reference this 2D
array through their methods. The different board status values (e.g. Empty, Hit)
correspond to integers for fast searching, comparison, and boolean methods.
*/
package board

import "errors"

// These two variables allow for conversion of each square status to
// the status string and vice-versa. This allows for statuses to be stored
// as integers for faster lookup and comparison.
var (
	values = map[string]int{
		"Empty": 0, "Miss": 1, "Destroyer": 2, "Submarine": 3,
		"Cruiser": 4, "Battleship": 5, "Carrier": 6, "Hit": 7,
	}

	status = [8]string{
		"Empty", "Miss", "Destroyer", "Submarine",
		"Cruiser", "Battleship", "Carrier", "Hit",
	}
)

// Board is a type for holding a standard 10x10 Battleship game board.
type Board [10][10]int

// Board update methods

// SetString sets a board value to a given string value.
func (b *Board) SetString(s Square, value string) error {
	if val, ok := values[value]; ok {
		b[s.Letter][s.Number] = val
		return nil
	}
	return errors.New("Given value is not a valid value")
}

// SetInt sets a board value to a given integer value.
func (b *Board) SetInt(s Square, value int) error {
	if value > 0 && value < 8 {
		b[s.Letter][s.Number] = value
		return nil
	}
	return errors.New("Given value out of range")
}

// SetPiece sets board values from a given piece's ship type and coordinates.
func (b *Board) SetPiece(p Piece) {
	value := values[p.Type.GetType()]
	for _, square := range p.Coords {
		b.SetInt(square, value)
	}
}

// PlacePiece sets board values from a given piece's ship type, but only if the
// squares are all empty. Useful when using Board for a player's ship placement,
// as opposed to tracking hits and misses.
func (b *Board) PlacePiece(p Piece) error {
	for _, square := range p.Coords {
		if !b.IsEmpty(square) {
			return errors.New("Piece coordinates are not empty")
		}
	}
	b.SetPiece(p)
	return nil
}

// Board retrieval methods

// GetString returns a given Square's string value.
func (b Board) GetString(s Square) string {
	return status[b[s.Letter][s.Number]]
}

// GetInt returns a given Square's integer value.
func (b Board) GetInt(s Square) int {
	return b[s.Letter][s.Number]
}

// Board boolean methods

// IsEmpty returns whether a given Square is empty.
func (b Board) IsEmpty(s Square) bool {
	return (b[s.Letter][s.Number] == 0)
}

// IsMiss returns whether a given Square is a miss.
func (b Board) IsMiss(s Square) bool {
	return (b[s.Letter][s.Number] == 1)
}

// IsHit returns whether a given square is any type of hit.
func (b Board) IsHit(s Square) bool {
	return (b[s.Letter][s.Number] > 1)
}

// IsUnsunk returns whether a given square is to an unsunk ship.
func (b Board) IsUnsunk(s Square) bool {
	return (b[s.Letter][s.Number] == 7)
}

// IsSunk returns whether a given square is to any sunk ship.
func (b Board) IsSunk(s Square) bool {
	val := b[s.Letter][s.Number]
	return (val > 1 && val < 7)
}

// IsShip returns whether a square belongs to a specific ship type.
func (b Board) IsShip(s Square, sh Ship) bool {
	return (status[b[s.Letter][s.Number]] == string(sh))
}
