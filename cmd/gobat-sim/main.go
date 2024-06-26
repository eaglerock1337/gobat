package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	gridSize    = 10
	grid        = [10][10]string{}
	currentView = "menu"
	lastClicked = ""
)

func initGrid() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid[i][j] = fmt.Sprintf("%c%d", 'A'+i, j+1)
		}
	}
}

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

	if v, err := g.SetView("side", maxX-20, 0, maxX-1, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Last Clicked"
		fmt.Fprintln(v, "None")
	}

	return nil
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
	viewName := v.Name()
	lastClicked = viewName
	sideView, err := g.View("side")
	if err != nil {
		return err
	}
	sideView.Clear()
	fmt.Fprintf(sideView, "Last clicked: %s\n", lastClicked)
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
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			viewName := grid[i][j]
			if err := g.SetKeybinding(viewName, gocui.MouseLeft, gocui.ModNone, mouseClick); err != nil {
				return err
			}
		}
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
	initGrid()

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
