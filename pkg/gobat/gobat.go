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
	"strconv"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/eaglerock1337/gobat/pkg/hunter"
	"github.com/jroimartin/gocui"
)

const (
	minX  = 71
	minY  = 31
	gridX = 50
)

var (
	h           *hunter.Hunter
	selectPos   = 1
	currentView = "menu"
	menuText    = []string{
		"H - Start Hunting",
		"R - Reset Hunter",
		"Q - Quit Gobat",
	}
)

// NewTerminal instantiates a gobat terminal screen
func NewTerminal() *gocui.Gui {
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

	hunt := hunter.NewHunter()
	hunt.Seek()
	h = &hunt

	return screen
}

// Run starts the main event loop of the application
func Run(g *gocui.Gui) {
	defer g.Close()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// RefreshSelectView refreshes the select view in the grid screen
func RefreshSelectView(v *gocui.View) {
	v.Clear()
	for i, square := range h.Shots {
		fmt.Fprintf(v, "%d - %s\n", i+1, square.PrintSquare())
	}
	v.SetCursor(0, selectPos)
	v.Highlight = false
	if v.Name() == currentView {
		v.Highlight = true
	}
}

// RefreshSquareView refreshes a specific square on the grid screen
func RefreshSquareView(v *gocui.View) {
	v.Clear()
	fmt.Fprintf(v, " %s\n", v.Name())
	square, _ := board.SquareByString(v.Name())
	if h.InShots(square) {
		v.BgColor = gocui.ColorGreen
	} else {
		v.BgColor = gocui.ColorDefault
	}
	fmt.Fprintf(v, " %d", h.HeatMap.GetSquare(square))
	v.SetCursor(0, 0)
	v.Highlight = false
	if v.Name() == currentView {
		v.Highlight = true
	}
}

// GridLayout provides the gocui manager function for the grid screen
func GridLayout(g *gocui.Gui) error {
	vertLine := "    |    |    |    |    |    |    |    |    |"
	horLine := "-------------------------------------------------"
	maxX, maxY := g.Size()

	if v, err := g.SetView("grid", 0, 0, gridX, minY-1); err != nil {
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

	letters := "ABCDEFGHIJ"
	for row := 0; row < 10; row++ {
		for col, letter := range letters {
			viewName := string(letter) + strconv.Itoa(row+1)
			if v, err := g.SetView(viewName, row*5, col*3, row*5+5, col*3+3); err != nil {
				if err != gocui.ErrUnknownView {
					return err
				}
				v.Frame = false
				v.SelBgColor = gocui.ColorWhite
				v.SelFgColor = gocui.ColorBlack
			} else {
				RefreshSquareView(v)
			}
		}
	}

	if maxX >= minX {
		if v, err := g.SetView("stats", gridX+1, 0, maxX-1, 2*minY/3); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Title = "Game Statistics"
			v.Wrap = true
		} else {
			v.Clear()
			perms := 0
			fmt.Fprintf(v, "Remaining ships:\n")
			for _, ship := range h.Ships {
				fmt.Fprintf(v, "  %s\n", ship.GetType())
				perms += h.Data[ship.GetLength()].Len()
			}
			fmt.Fprintf(v, "\nTurns Taken: %d\n", h.Turns)
			fmt.Fprintf(v, "Permutations: %d\n", perms)
			fmt.Fprintf(v, "\nActive Hitstack:\n")
			for _, square := range h.HitStack {
				fmt.Fprintf(v, "%s ", square.PrintSquare())
			}
			if len(h.HitStack) == 0 {
				fmt.Fprint(v, "  Empty")
			}
		}

		if v, err := g.SetView("select", gridX+1, 2*minY/3+1, maxX-1, minY-1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Title = "Select Option"
			v.Wrap = true
			v.Highlight = true
			v.SelBgColor = gocui.ColorWhite
			v.SelFgColor = gocui.ColorBlack
			RefreshSelectView(v)
		}

		if _, err := g.SetCurrentView(currentView); err != nil {
			return err
		}
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
	if v, err := g.SetView("menubg", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	} else {
		if maxX < minX || maxY < minY {
			v.BgColor = gocui.ColorRed
		} else {
			v.BgColor = gocui.ColorDefault
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
	currentView = "select"
	maxX, maxY := g.Size()
	if minX <= maxX && minY <= maxY {
		g.SetManagerFunc(GridLayout)
		SetKeyBindings(g)
		selectPos = 0
	}
	return nil
}

// SwitchToMenu switches to menu view
func SwitchToMenu(g *gocui.Gui, v *gocui.View) error {
	currentView = "menu"
	g.SetManagerFunc(MenuLayout)
	SetKeyBindings(g)
	return nil
}

// SelectCursorDown handles the cursor down keybind in the select view
func SelectCursorDown(g *gocui.Gui, v *gocui.View) error {
	selectPos++
	if selectPos > 4 {
		selectPos = 4
	}
	RefreshSelectView(v)
	return nil
}

// SelectCursorUp handles the cursor down keybind in the select view
func SelectCursorUp(g *gocui.Gui, v *gocui.View) error {
	selectPos--
	if selectPos < 0 {
		selectPos = 0
	}
	RefreshSelectView(v)
	return nil
}

// SelectCursorLeft handles the cursor left keybind in the select view
func SelectCursorLeft(g *gocui.Gui, v *gocui.View) error {
	currentView = "J10"
	if _, err := g.SetCurrentView("J10"); err != nil {
		return err
	}
	RefreshSelectView(v)
	RefreshSquareView(g.CurrentView())
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

	if err := g.SetKeybinding("select", gocui.KeyArrowDown, gocui.ModNone, SelectCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("select", gocui.KeyArrowUp, gocui.ModNone, SelectCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("select", gocui.KeyArrowLeft, gocui.ModNone, SelectCursorLeft); err != nil {
		return err
	}
	// if err := g.SetKeybinding("select", gocui.KeyEnter, gocui.ModNone, SelectOption); err != nil {
	// 	return err
	// }
	return nil
}
