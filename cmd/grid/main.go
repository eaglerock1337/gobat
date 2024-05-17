package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	gridSize    = 10
	grid        = [10][10]string{}
	lastClicked = ""
)

func initGrid() {
	// Initialize the grid with square names
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid[i][j] = fmt.Sprintf("%c%d", 'A'+i, j+1)
		}
	}
}

func layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	// Create the grid views
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			viewName := grid[i][j]
			v, err := g.SetView(viewName, j*4, i*2, j*4+3, i*2+1)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
			v.Title = viewName
		}
	}

	// Create the side view
	if v, err := g.SetView("side", maxX-20, 0, maxX-1, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Last Clicked"
		fmt.Fprintln(v, "None")
	}

	return nil
}

func mouseClick(g *gocui.Gui, v *gocui.View) error {
	viewName := v.Name()

	// Update the side view with the last clicked square
	lastClicked = viewName
	sideView, err := g.View("side")
	if err != nil {
		return err
	}
	sideView.Clear()
	fmt.Fprintf(sideView, "Last clicked: %s\n", lastClicked)

	return nil
}

func keybindings(g *gocui.Gui) error {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			viewName := grid[i][j]
			if err := g.SetKeybinding(viewName, gocui.MouseLeft, gocui.ModNone, mouseClick); err != nil {
				return err
			}
		}
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	return nil
}

func main() {
	initGrid()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
