package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

const (
	gridSize     = 10
	squareWidth  = 4
	squareHeight = 2
)

var (
	currentView = "menu"
	lastClicked = ""
)

func layout(g *gocui.Gui) error {
	if currentView == "menu" {
		return layoutMenu(g)
	}
	return layoutGrid(g)
}

func layoutMenu(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("menu", maxX/4, maxY/4, 3*maxX/4, 3*maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main Menu"
		fmt.Fprintln(v, "1. Start Game")
		fmt.Fprintln(v, "Q. Quit")
		if _, err := g.SetCurrentView("menu"); err != nil {
			return err
		}
	}
	return nil
}

func layoutGrid(g *gocui.Gui) error {
	maxX, _ := g.Size()
	gridViewWidth := gridSize * squareWidth
	gridViewHeight := gridSize * squareHeight

	if v, err := g.SetView("grid", 0, 0, gridViewWidth+1, gridViewHeight+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Battleship Grid"
		drawGrid(v)
	}

	if v, err := g.SetView("side", gridViewWidth+2, 0, maxX-1, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Last Clicked"
		fmt.Fprintln(v, "None")
	}

	return nil
}

func drawGrid(v *gocui.View) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			cell := fmt.Sprintf("%c%d", 'A'+i, j+1)
			fmt.Fprintf(v, "%s", cell)
			for k := 0; k < squareHeight-1; k++ {
				fmt.Fprint(v, "\n")
			}
			if j < gridSize-1 {
				fmt.Fprint(v, " ")
			}
		}
		if i < gridSize-1 {
			fmt.Fprint(v, "\n")
		}
	}
}

func switchToGrid(g *gocui.Gui, v *gocui.View) error {
	currentView = "grid"
	g.SetManagerFunc(layout)
	updateKeybindings(g)
	return nil
}

func switchToMenu(g *gocui.Gui, v *gocui.View) error {
	currentView = "menu"
	g.SetManagerFunc(layout)
	updateKeybindings(g)
	return nil
}

func mouseClick(g *gocui.Gui, v *gocui.View) error {
	if v.Name() != "grid" {
		return nil
	}

	x, y := v.Cursor()
	col := x / squareWidth
	row := y / squareHeight
	if col < gridSize && row < gridSize {
		lastClicked = fmt.Sprintf("%c%d", 'A'+row, col+1)
		sideView, err := g.View("side")
		if err != nil {
			return err
		}
		sideView.Clear()
		fmt.Fprintf(sideView, "Last clicked: %s\n", lastClicked)
	}
	return nil
}

func keybindingsMenu(g *gocui.Gui) error {
	if err := g.SetKeybinding("menu", '1', gocui.ModNone, switchToGrid); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", 'q', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}
	return nil
}

func keybindingsGrid(g *gocui.Gui) error {
	if err := g.SetKeybinding("grid", gocui.MouseLeft, gocui.ModNone, mouseClick); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'm', gocui.ModNone, switchToMenu); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}
	return nil
}

func updateKeybindings(g *gocui.Gui) error {
	g.DeleteKeybindings("")
	if currentView == "menu" {
		return keybindingsMenu(g)
	}
	return keybindingsGrid(g)
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := updateKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
