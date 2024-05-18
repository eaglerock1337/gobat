package gobat

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var menuText = []string{
	"H - Start Hunting",
	"R - Reset Hunter",
	"Q - Quit Gobat",
}

// menuLayout provides the gocui manager function for the main menu
func menuLayout(g *gocui.Gui) error {
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

// switchToMenu switches to menu view
func switchToMenu(g *gocui.Gui, v *gocui.View) error {
	currentView = "menu"
	g.SetManagerFunc(menuLayout)
	setKeyBindings(g)
	selectPos = 0
	return nil
}

// menuSelection handles the selection of a menu item
func menuSelection(g *gocui.Gui, v *gocui.View) {
	currentView = v.Name()
}
