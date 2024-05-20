package gobat

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var (
	promptName      string
	promptOptions   []string
	promptSelection int
)

// showPromptView shows the general prompt view
func showPromptView(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("prompt", maxX/2-7, maxY/2-4, maxX/2+7, maxY/2+4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
	} else {
		refreshPromptView(v)
	}

	return nil
}

// refreshPromptView refreshes the general prompt view
func refreshPromptView(v *gocui.View) {
	v.Clear()
	for _, line := range promptOptions {
		fmt.Fprintln(v, line)
	}
	v.SetCursor(0, promptSelection)

	v.Title = promptName
	v.Highlight = false
	if currentView == "select" {
		v.Highlight = true
	}
}

// promptEnterKeySelection processes any enter key prompt selection
func promptEnterKeySelection(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "menu", "menubg", "error":
		return nil
	default:
		if err := gridPromptEnterKeySelection(g, v); err != nil {
			return err
		}
	}
	return nil
}

// if condition {
// 	if v, err := g.SetViewOnTop("error"); err == nil {
// 		do stuff
// 	} else {
// 		return err
// 	}
// } else {
// 	if _, err := g.SetViewOnBottom("error"); err != nil {
// 		return err
// 	}
// }
