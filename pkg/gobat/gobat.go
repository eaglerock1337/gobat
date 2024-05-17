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
	"time"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/eaglerock1337/gobat/pkg/hunter"
	"github.com/jroimartin/gocui"
)

var (
	minX     = 71
	minY     = 31
	menuText = []string{
		"H - Start Hunting",
		"R - Reset Hunter",
		"Q - Quit Gobat",
	}
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
	screen.SetCurrentView("menu")
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
	maxX, maxY := g.Size()

	if v, err := g.SetView("grid", 0, 0, 50, 30); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Battleship Grid"
		for i := 1; i <= 30; i++ {
			line := vertLine
			if i%3 == 0 {
				line = horLine
			}
			fmt.Fprintln(v, line)
		}
	}

	if v, err := g.SetView("side", 51, 0, 70, 30); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	} else {
		v.Title = "Controls"
		v.Wrap = true
		v.Clear()
		fmt.Fprintln(v, "Side view stuff go here.")
	}

	if _, err := g.SetCurrentView("side"); err != nil {
		return err
	}

	if v, err := g.SetView("error", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Screen too small"
	} else {
		if maxX < minX || maxY < minY {
			v.BgColor = gocui.ColorRed
			if v, err := g.SetViewOnTop("error"); err == nil {
				v.Clear()
				fmt.Fprintf(v, "Min Size: %dx%d\n", minX, minY)
				fmt.Fprintf(v, "Cur Size: %dx%d\n", maxX, maxY)
				fmt.Fprintln(v, "M - Main Menu")
				fmt.Fprintln(v, "Q - Quit")
			} else {
				return err
			}
		} else {
			if _, err := g.SetViewOnBottom("error"); err != nil {
				return err
			}
		}
	}

	return nil
}

// MenuLayout provides the gocui manager function for the main menu
func MenuLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("menubg", 0, 0, maxX-1, maxY-1); err == nil {
		if maxX < minX || maxY < minY {
			v.BgColor = gocui.ColorRed
		} else {
			v.BgColor = gocui.ColorDefault
		}
	} else {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if v, err := g.SetView("menu", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Gobat Hunter"
	} else {
		v.Clear()
		for _, line := range menuText {
			fmt.Fprintln(v, line)
		}
		fmt.Fprintf(v, "\nMin Size: %dx%d", minX, minY)
		fmt.Fprintf(v, "\nCur Size: %dx%d", maxX, maxY)
		g.SetCurrentView("menu")
	}

	return nil
}

// Quit terminates the screen and event loop
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// SwitchToGrid switches to grid view
func SwitchToGrid(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if maxX < minX || maxY < minY {
		v.BgColor = gocui.ColorRed
		time.Sleep(time.Second / 10)
		v.BgColor = gocui.ColorDefault
	} else {
		g.SetManagerFunc(GridLayout)
		SetKeyBindings(g)
	}
	return nil
}

// SwitchToMenu switches to menu view
func SwitchToMenu(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(MenuLayout)
	SetKeyBindings(g)
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
	if err := g.SetKeybinding("", 'm', gocui.ModNone, SwitchToMenu); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", 'h', gocui.ModNone, SwitchToGrid); err != nil {
		return err
	}

	return nil
}
