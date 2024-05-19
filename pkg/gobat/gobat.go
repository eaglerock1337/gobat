/*
Package gobat is responsible for creating and managing the console display for
the gobat CLI application. It is responsible for creating an ncurses-based window
that will be used for displaying the game board and top moves.
*/

package gobat

import (
	"fmt"
	"log"

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
	currentView = "menu"
)

// NewTerminal instantiates a gobat terminal screen
func NewTerminal() *gocui.Gui {
	screen, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	screen.SetManagerFunc(menuLayout)
	screen.SetCurrentView(currentView)
	screen.Mouse = true

	if err := setKeyBindings(screen); err != nil {
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

// quit terminates the screen and event loop
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// enterKey handles enter key input
func enterKey(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		g.SetCurrentView(currentView)
		v = g.CurrentView()
	}
	switch v.Name() {
	case "menu", "menubg":
		if err := menuEnterKeySelection(g, v); err != nil {
			return err
		}
	default:
		if err := gridEnterKeySelection(g, v); err != nil {
			return err
		}
	}
	return nil
}

// mouseClick handles mouse click input
func mouseClick(g *gocui.Gui, v *gocui.View) error {
	switch v.Name() {
	case "error", "grid", "menubg", "stats":
		return nil
	case "menu":
		menuMouseClickSelection(g, v)
	default:
		gridMouseClickSelection(g, v)
	}
	return nil
}

// setKeyBindings sets all gocui keybindings
func setKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'm', gocui.ModNone, switchToMenu); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'g', gocui.ModNone, switchToGrid); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, cursorLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, cursorRight); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, enterKey); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, mouseClick); err != nil {
		return err
	}
	return nil
}

// initializePromptView initializes the general prompt view
func initializePromptView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("prompt", gridX+1, 2*minY/3+1, maxX-1, minY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Select Option"
		v.Wrap = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
	} else {
		refreshSelectView(v)
	}

	return nil
}

// refreshPromptView refreshes the select view in the grid screen
func refreshPromptView(v *gocui.View) {
	v.Clear()

	for i, square := range h.Shots {
		fmt.Fprintf(v, "%d - %s (%d)\n", i+1, square.PrintSquare(), h.HeatMap.GetSquare(square))
	}
	for _, line := range gridControls {
		fmt.Fprintln(v, line)
	}
	v.SetCursor(0, gridSelection)

	v.Highlight = false
	if currentView == "select" {
		v.Highlight = true
	}
}
