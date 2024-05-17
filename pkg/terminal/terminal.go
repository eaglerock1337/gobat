/*
Package terminal is responsible for creating and managing the console display for
the gobat CLI application. It is responsible for creating an ncurses-based window
that will be used for displaying the game board and top moves. The display struct
contains the initialized tcell ncurses window as well as necessary member
functions for interfacing with the ncurses environment during use of the program.
*/

package terminal

import (
	"errors"
	"fmt"
	"log"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/eaglerock1337/gobat/pkg/hunter"
	"github.com/jroimartin/gocui"
)

// Display is a custom type of tcell.Screen with custom gobat-related methods
type Terminal struct {
	Screen *gocui.Gui // a gocui
}

// MenuLayout provides the gocui manager function for the main menu
func MenuLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello, Gocui!")
	}
	return nil
}

// Quit terminates the screen and event loop
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// NewTerminal instantiates a terminal struct including a gocui screen
func NewTerminal() Terminal {
	scr, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	scr.SetManagerFunc(MenuLayout)

	if err := scr.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	if err := scr.SetKeybinding("", 'q', gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	newTerminal := Terminal{scr}
	return newTerminal
}

// Run starts the main event loop of the application
func (d Terminal) Run() {
	fmt.Printf("d: %v\n", d)
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
