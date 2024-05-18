package gobat

import (
	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/jroimartin/gocui"
)

// cursorDown handles the gocui cursor down keybind
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "error", "grid", "menu", "menubg", "stats":
		return nil
	case "select":
		selectPos++
		if selectPos > 4 {
			selectPos = 4
		}
		refreshSelectView(v)
	default:
		curSquare, _ := board.SquareByString(currentView)
		newSquare, err := board.SquareByValue(curSquare.Letter, curSquare.Number+1)
		if err != nil {
			return nil
		}
		currentView = newSquare.PrintSquare()
		if _, err := g.SetCurrentView(currentView); err != nil {
			return err
		}
	}
	return nil
}

// cursorUp handles the gocui cursor down keybind
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "error", "grid", "menu", "menubg", "stats":
		return nil
	case "select":
		selectPos--
		if selectPos < 0 {
			selectPos = 0
		}
		refreshSelectView(v)
	default:
		curSquare, _ := board.SquareByString(currentView)
		newSquare, err := board.SquareByValue(curSquare.Letter, curSquare.Number-1)
		if err != nil {
			return nil
		}
		currentView = newSquare.PrintSquare()
		if _, err := g.SetCurrentView(currentView); err != nil {
			return err
		}
	}
	return nil
}

// cursorLeft handles the gocui cursor left keybind
func cursorLeft(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "select":
		currentView = "J10"
		if _, err := g.SetCurrentView("J10"); err != nil {
			return err
		}
		refreshSelectView(v)
	case "error", "grid", "menu", "menubg", "stats":
		return nil
	default:
		curSquare, _ := board.SquareByString(currentView)
		if curSquare.Letter > 0 {
			newSquare, err := board.SquareByValue(curSquare.Letter-1, curSquare.Number)
			if err != nil {
				return err
			}
			currentView = newSquare.PrintSquare()
			if _, err := g.SetCurrentView(currentView); err != nil {
				return err
			}
		}
	}
	refreshSquareView(g.CurrentView())
	return nil
}

// cursorRight handles the gocui cursor Right keybind
func cursorRight(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "error", "grid", "menu", "menubg", "select", "stats":
		return nil
	default:
		curSquare, _ := board.SquareByString(currentView)
		if curSquare.Letter == 9 {
			currentView = "select"
			selectPos = 0
		} else {
			newSquare, err := board.SquareByValue(curSquare.Letter+1, curSquare.Number)
			if err != nil {
				return err
			}
			currentView = newSquare.PrintSquare()
		}
		if _, err := g.SetCurrentView(currentView); err != nil {
			return err
		}
		refreshSquareView(v)
		if currentView == "select" {
			refreshSelectView(g.CurrentView())
		} else {
			refreshSquareView(g.CurrentView())
		}
	}
	return nil
}
