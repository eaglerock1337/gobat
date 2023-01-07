/*
Package terminal is responsible for creating and managing the console display for
the gobat CLI application. It is responsible for creating an ncurses-based window
that will be used for displaying the game board and top moves. The display struct
contains the initialized tcell ncurses window as well as necessary member
functions for interfacing with the ncurses environment during use of the program.
*/

package display

import (
	"errors"
	"fmt"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/eaglerock1337/gobat/pkg/hunter"
	"github.com/gdamore/tcell/v2"
)

// Display is a custom type of tcell.Screen with custom gobat-related methods
type Display tcell.Screen

func NewDisplay() Display {
	var newDisplay Display

	return newDisplay
}

// This is literally just here to stop Go from deleting imports I'll need later
func AnnoyingErrors() (error, error) {
	square, err := board.SquareByString("A1")
	if err != nil {
		niceErr := fmt.Errorf("this is a formatted error: %v", err)
		return errors.New("this is an error"), niceErr
	}

	hunter := hunter.NewHunter()
	hunter.AddShot(square)
	hunter.ClearShots()
	return nil, nil
}
