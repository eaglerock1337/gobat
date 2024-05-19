package gobat

import (
	"fmt"
	"strconv"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/jroimartin/gocui"
)

var gridSelection = 0

var selectedSquare board.Square

var gridControls = []string{
	// "H - Help",
	"M - Menu",
	"Q - Quit",
}

// gridLayout provides the gocui manager function for the grid screen
func gridLayout(g *gocui.Gui) error {
	if err := showGridView(g); err != nil {
		return err
	}
	if err := showSquareViews(g); err != nil {
		return err
	}
	if err := showSideViews(g); err != nil {
		return err
	}
	if err := showErrorView(g); err != nil {
		return err
	}
	return nil
}

// gridEnterKeySelection handles enter key selection on any grid square
func gridEnterKeySelection(g *gocui.Gui, v *gocui.View) error {
	switch currentView {
	case "select", "stats":
		if gridSelection < len(h.Shots) {
			if err := h.Turn(h.Shots[gridSelection], "Miss"); err != nil {
				return err
			}
		}
	}
	switchToGrid(g, v)
	return nil
}

// gridMouseClickSelection handles mouse click selection of a specific grid square
func gridMouseClickSelection(g *gocui.Gui, v *gocui.View) {
	n := v.Name()
	g.SetCurrentView(n)
}

// showGridView shows the grid view in the grid screen
func showGridView(g *gocui.Gui) error {
	vertLine := "    |    |    |    |    |    |    |    |    |"
	horLine := "-------------------------------------------------"

	if v, err := g.SetView("grid", 0, 0, gridX, minY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Battleship Grid"
		for i := 1; i <= 30; i++ {
			line := vertLine
			if i%3 == 0 {
				line = horLine
			}
			fmt.Fprintln(v, line)
		}
	}

	return nil
}

// showSquareViews shows all square views in the grid screen
func showSquareViews(g *gocui.Gui) error {
	letters := "ABCDEFGHIJ"

	for row, letter := range letters {
		for col := 0; col < 10; col++ {
			viewName := string(letter) + strconv.Itoa(col+1)
			if v, err := g.SetView(viewName, row*5, col*3, row*5+5, col*3+3); err != nil {
				if err != gocui.ErrUnknownView {
					return err
				}
				v.Frame = false
				v.SelBgColor = gocui.ColorWhite
				v.SelFgColor = gocui.ColorBlack
			} else {
				refreshSquareView(v)
			}
		}
	}

	return nil
}

// refreshSquareView refreshes a specific square on the grid screen
func refreshSquareView(v *gocui.View) {
	v.Clear()

	fmt.Fprintf(v, " %s \n", v.Name())
	square, _ := board.SquareByString(v.Name())
	fmt.Fprintf(v, " %d", h.HeatMap.GetSquare(square))
	v.SetCursor(0, 0)

	if h.InShots(square) {
		v.BgColor = gocui.ColorGreen
	} else {
		v.BgColor = gocui.ColorDefault
	}

	v.Highlight = false
	if v.Name() == currentView {
		v.Highlight = true
	}
}

// showSideViews shows all side views in the grid screen
func showSideViews(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if maxX < minX {
		return nil
	}
	if err := showStatsView(g); err != nil {
		return err
	}
	if err := showSelectView(g); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(currentView); err != nil {
		return err
	}
	return nil
}

// showStatsView shows the stats view in the grid screen
func showStatsView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("stats", gridX+1, 0, maxX-1, 2*minY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Game Statistics"
		v.Wrap = true
	} else {
		refreshStatsView(v)
	}

	return nil
}

// refreshStatsView refreshes the status view on the grid screen
func refreshStatsView(v *gocui.View) {
	v.Clear()

	perms := 0
	fmt.Fprintf(v, "Remaining ships:\n")

	for _, ship := range h.Ships {
		fmt.Fprintf(v, "  %s\n", ship.GetType())
		perms += h.Data[ship.GetLength()].Len()
	}

	fmt.Fprintf(v, "\nTurns Taken: %d\n", h.Turns)
	fmt.Fprintf(v, "Total Perms: %d\n", perms)

	mode := "Destroy"
	if h.SeekMode {
		mode = "Seek"
	}
	fmt.Fprintf(v, "Hunter: %s\n", mode)

	fmt.Fprintf(v, "\nActive Hitstack:\n")
	for _, square := range h.HitStack {
		fmt.Fprintf(v, "%s ", square.PrintSquare())
	}
	if len(h.HitStack) == 0 {
		fmt.Fprint(v, "  Empty")
	}
}

// showSelectView shows the select view in the grid screen
func showSelectView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("select", gridX+1, 2*minY/3+1, maxX-1, minY-1); err != nil {
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

// refreshSelectView refreshes the select view in the grid screen
func refreshSelectView(v *gocui.View) {
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

// showErrorView shows the error view
func showErrorView(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("error", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Screen too small"
	} else {
		if err := refreshErrorView(g, v); err != nil {
			return err
		}
	}
	return nil
}

// refreshErrorView refreshes the error view in the grid screen
func refreshErrorView(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if maxX < minX || maxY < minY {
		v.BgColor = gocui.ColorRed
		if v, err := g.SetViewOnTop("error"); err == nil {
			v.Clear()
			fmt.Fprintf(v, "Need: %dx%d\n", minX, minY)
			fmt.Fprintf(v, "Have: %dx%d\n", maxX, maxY)
			for _, line := range gridControls {
				fmt.Fprintln(v, line)
			}
		} else {
			return err
		}
	} else {
		if _, err := g.SetViewOnBottom("error"); err != nil {
			return err
		}
	}

	return nil
}

// switchToGrid switches to grid view
func switchToGrid(g *gocui.Gui, v *gocui.View) error {
	currentView = "select"
	maxX, maxY := g.Size()

	if minX <= maxX && minY <= maxY {
		gridSelection = 0
		g.SetCurrentView(currentView)

		g.SetManagerFunc(gridLayout)
		setKeyBindings(g)
	}

	return nil
}
