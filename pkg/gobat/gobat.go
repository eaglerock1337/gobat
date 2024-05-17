/*
Package gobat is responsible for creating and managing the console display for
the gobat CLI application. It is responsible for creating an ncurses-based window
that will be used for displaying the game board and top moves. The display struct
contains the initialized tcell ncurses window as well as necessary member
functions for interfacing with the ncurses environment during use of the program.
*/

package gobat

import (
	"fmt"
	"log"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/eaglerock1337/gobat/pkg/hunter"
	"github.com/jroimartin/gocui"
)

// Gobat provides all data required for interactive hunter
type Gobat struct {
	Screen *gocui.Gui    // gocui terminal screen
	Hunter hunter.Hunter // ideal strategy hunter
}

// NewGobat instantiates a gobat struct including a gocui screen and hunter
func NewGobat() Gobat {
	screen, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	screen.SetManagerFunc(MenuLayout)
	screen.Mouse = true

	if err := SetKeyBindings(screen); err != nil {
		log.Panicln(err)
	}

	hunter := hunter.NewHunter()
	newGobat := Gobat{screen, hunter}
	return newGobat
}

// Run starts the main event loop of the application
func (g Gobat) Run() {
	defer g.Screen.Close()

	if err := g.Screen.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// This is literally just here to stop Go from deleting imports I'll need later
func AnnoyingErrors() {
	square, _ := board.SquareByString("A1")
	log.Println(square.PrintSquare())
}

// GridLayout provides the gocui manager function for the grid menu
func GridLayout(g *gocui.Gui) error {
	vertLine := "    |    |    |    |    |    |    |    |    |"
	horLine := "-------------------------------------------------"

	if v, err := g.SetView("grid", 0, 50, 0, 30); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Battleship Grid"
		for i := 0; i < 30; i++ {
			line := vertLine
			if i%3 == 0 {
				line = horLine
			}
			fmt.Fprintln(v, line)
		}
	}
	return nil
}

// MenuLayout provides the gocui manager function for the main menu
func MenuLayout(g *gocui.Gui) error {
	minX, minY := 71, 31
	maxX, maxY := g.Size()
	if v, err := g.SetView("menu", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = fmt.Sprintf("Gobat Hunter - %dx%d", maxX, maxY)
		menuText := []string{
			"H - Start Hunting",
			"R - Reset Hunter",
			"Q - Quit Gobat",
			"",
		}
		for _, line := range menuText {
			fmt.Fprintln(v, line)
		}
		fmt.Fprintf(v, "Min Size: %dx%d\n", minX, minY)
	} else {
		v.Highlight = false
		if maxX < minX || maxY < minY {
			v.BgColor = gocui.ColorRed
		} else {
			v.BgColor = gocui.ColorDefault
		}
		v.Title = fmt.Sprintf("Gobat Hunter - %dx%d", maxX, maxY)
	}
	return nil
}

// Quit terminates the screen and event loop
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// SwitchToGrid switches to grid view
func SwitchToGrid(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(GridLayout)
	return nil
}

// SwitchToMenu switches to menu view
func SwitchToMenu(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(MenuLayout)
	return nil
}

// SetKeyBindings sets all global and view keybindings
func SetKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, Quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", 'h', gocui.ModNone, SwitchToGrid); err != nil {
		return err
	}
	if err := g.SetKeybinding("grid", 'm', gocui.ModNone, SwitchToMenu); err != nil {
		return err
	}
	return nil
}
