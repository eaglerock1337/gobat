package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

type Menu struct {
	Title string
	Items []string
}

var (
	menus = map[string]Menu{
		"main": {
			Title: "Main Menu",
			Items: []string{"A. Go to Submenu A", "B. Go to Submenu B", "C. Option C", "Q. Quit"},
		},
		"submenuA": {
			Title: "Submenu A",
			Items: []string{"1. Suboption A1", "2. Suboption A2", "B. Back to Main Menu"},
		},
		"submenuB": {
			Title: "Submenu B",
			Items: []string{"1. Suboption B1", "2. Suboption B2", "B. Back to Main Menu"},
		},
	}
	currentMenu = "main"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("menu", maxX/4, maxY/4, 3*maxX/4, 3*maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		updateMenuView(v)
		if _, err := g.SetCurrentView("menu"); err != nil {
			return err
		}
	}
	return nil
}

func updateMenuView(v *gocui.View) {
	v.Clear()
	menu := menus[currentMenu]
	v.Title = menu.Title
	for _, item := range menu.Items {
		fmt.Fprintln(v, item)
	}
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if cy < len(menus[currentMenu].Items)-1 {
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
	item := menus[currentMenu].Items[cy]
	return handleSelection(g, item)
}

func handleSelection(g *gocui.Gui, item string) error {
	switch currentMenu {
	case "main":
		switch item {
		case "A. Go to Submenu A":
			currentMenu = "submenuA"
		case "B. Go to Submenu B":
			currentMenu = "submenuB"
		case "C. Option C":
			fmt.Println("Option C selected")
		case "Q. Quit":
			return gocui.ErrQuit
		}
	case "submenuA":
		switch item {
		case "1. Suboption A1":
			fmt.Println("Suboption A1 selected")
		case "2. Suboption A2":
			fmt.Println("Suboption A2 selected")
		case "B. Back to Main Menu":
			currentMenu = "main"
		}
	case "submenuB":
		switch item {
		case "1. Suboption B1":
			fmt.Println("Suboption B1 selected")
		case "2. Suboption B2":
			fmt.Println("Suboption B2 selected")
		case "B. Back to Main Menu":
			currentMenu = "main"
		}
	}
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("menu")
		if err != nil {
			return err
		}
		updateMenuView(v)
		return nil
	})
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
