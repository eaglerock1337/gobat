package gobat

import (
	"fmt"
	"strconv"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/jroimartin/gocui"
)

// gridLayout provides the gocui manager function for the grid screen
func gridLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := initializeGridView(g); err != nil {
		return err
	}

	if err := initializeSquareViews(g); err != nil {
		return err
	}

	if err := initializeSideViews(g); err != nil {
		return err
	}

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

// gridSelection handles the selection of a specific grid square
func gridSelection(g *gocui.Gui, v *gocui.View) {
	n := v.Name()
	g.SetCurrentView(n)
}

// initializeGridView initializes the grid view in the grid screen
func initializeGridView(g *gocui.Gui) error {
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

// initializeSquareViews initializes all square views in the grid screen
func initializeSquareViews(g *gocui.Gui) error {
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

// initializeStatsView initializes the stats view in the grid screen
func initializeStatsView(g *gocui.Gui) error {
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

// initializeSelectView initializes the select view in the grid screen
func initializeSelectView(g *gocui.Gui) error {
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

// initializeSideViews initializes all side views in the grid screen
func initializeSideViews(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if maxX < minX {
		return nil
	}
	if err := initializeStatsView(g); err != nil {
		return err
	}
	if err := initializeSelectView(g); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(currentView); err != nil {
		return err
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
			fmt.Fprintf(v, "Min Size: %dx%d\n", minX, minY)
			fmt.Fprintf(v, "Cur Size: %dx%d\n", maxX, maxY)
			fmt.Fprintln(v, "M - Main Menu")
			fmt.Fprintln(v, "Q - Quit")
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

// refreshSelectView refreshes the select view in the grid screen
func refreshSelectView(v *gocui.View) {
	v.Clear()

	for i, square := range h.Shots {
		fmt.Fprintf(v, "%d - %s\n", i+1, square.PrintSquare())
	}
	v.SetCursor(0, selectPos)

	v.Highlight = false
	if currentView == "select" {
		v.Highlight = true
	}
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
	fmt.Fprintf(v, "Permutations: %d\n", perms)

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

	fmt.Fprintf(v, "\n\nH - Help")
	fmt.Fprintf(v, "\nM - Menu")
	fmt.Fprintf(v, "\nQ - Quit")
}

// switchToGrid switches to grid view
func switchToGrid(g *gocui.Gui, v *gocui.View) error {
	currentView = "select"
	maxX, maxY := g.Size()

	if minX <= maxX && minY <= maxY {
		g.SetManagerFunc(gridLayout)
		setKeyBindings(g)
		selectPos = 0
	}

	return nil
}
