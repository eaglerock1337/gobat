package gobat

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var menuSelection = 0
var menuControls = []string{
	"G - Go Hunting",
	// "R - Reset Hunter",
	"Q - Quit Gobat",
}

// menuLayout provides the gocui manager function for the main menu
func menuLayout(g *gocui.Gui) error {
	if err := initializeMenuBackgroundView(g); err != nil {
		return err
	}
	if err := initializeMenuView(g); err != nil {
		return err
	}
	return nil
}

// menuEnterKeySelection handles menu enter key selection
func menuEnterKeySelection(g *gocui.Gui, v *gocui.View) error {
	switch menuSelection {
	case 0:
		if err := switchToGrid(g, v); err != nil {
			return err
		}
	case 1:
		g.Update(func(g *gocui.Gui) error {
			return gocui.ErrQuit
		})
	}
	return nil
}

// menuMouseClickSelection handles menu mouse click selection
func menuMouseClickSelection(g *gocui.Gui, v *gocui.View) {
	currentView = v.Name()
	g.SetCurrentView(currentView)
}

// initializeMenuView initializes the menu view in the menu screen
func initializeMenuView(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("menu", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Gobat Hunter"
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
	} else {
		refreshMenuView(g, v)
	}

	return nil
}

// refreshMenuView refreshes the menu view in the menu screen
func refreshMenuView(g *gocui.Gui, v *gocui.View) {
	maxX, maxY := g.Size()
	v.Clear()
	for _, line := range menuControls {
		fmt.Fprintln(v, line)
	}
	fmt.Fprintf(v, "\nMin Size: %dx%d", minX, minY)
	fmt.Fprintf(v, "\nCur Size: %dx%d", maxX, maxY)

	v.SetCursor(0, menuSelection)
	v.SelBgColor = gocui.ColorWhite
	if menuSelection == 0 && (maxX < minX || maxY < minY) {
		v.SelBgColor = gocui.ColorRed
	}
}

// initializeMenuBackgroundView initializes the background view in the menu screen
func initializeMenuBackgroundView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("menubg", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

// switchToMenu switches to menu view
func switchToMenu(g *gocui.Gui, v *gocui.View) error {
	if currentView == "menu" {
		return nil
	}
	currentView = "menu"
	g.SetManagerFunc(menuLayout)
	setKeyBindings(g)
	menuSelection = 0
	refreshMenuView(g, v)
	return nil
}
