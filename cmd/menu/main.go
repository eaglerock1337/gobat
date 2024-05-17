package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var menuItems = []string{
	"A. Option A",
	"B. Option B",
	"C. Option C",
	"Q. Quit",
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("menu", maxX/4, maxY/4, 3*maxX/4, 3*maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main Menu"
		for _, item := range menuItems {
			fmt.Fprintln(v, item)
		}
		if _, err := g.SetCurrentView("menu"); err != nil {
			return err
		}
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if cy < len(menuItems)-1 {
		v.SetCursor(cx, cy+1)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if cy > 0 {
		v.SetCursor(0, cy-1)
	}
	return nil
}

func enter(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	item := menuItems[cy]
	return handleSelection(g, item)
}

func handleSelection(g *gocui.Gui, item string) error {
	switch item {
	case "A. Option A":
		fmt.Println("Option A selected")
	case "B. Option B":
		fmt.Println("Option B selected")
	case "C. Option C":
		fmt.Println("Option C selected")
	case "Q. Quit":
		return gocui.ErrQuit
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, enter); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'a', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return handleSelection(g, "A. Option A")
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'b', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return handleSelection(g, "B. Option B")
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'c', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return handleSelection(g, "C. Option C")
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}
	return nil
}

func main() {
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
