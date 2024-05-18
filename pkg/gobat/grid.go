package gobat

import (
	"fmt"
	"strconv"

	"github.com/eaglerock1337/gobat/pkg/board"
	"github.com/jroimartin/gocui"
)

// GridLayout provides the gocui manager function for the grid screen
func GridLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := InitializeGridView(g); err != nil {
		return err
	}

	if err := InitializeSquareViews(g); err != nil {
		return err
	}

	if err := InitializeSideViews(g); err != nil {
		return err
	}

	if v, err := g.SetView("error", maxX/3, maxY/3, 2*maxX/3, 2*maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Screen too small"
	} else {
		if err := RefreshErrorView(g, v); err != nil {
			return err
		}
	}

	return nil
}

// InitializeGridView initializes the grid view in the grid screen
func InitializeGridView(g *gocui.Gui) error {
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

// InitializeSquareViews initializes all square views in the grid screen
func InitializeSquareViews(g *gocui.Gui) error {
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
				RefreshSquareView(v)
			}
		}
	}

	return nil
}

// InitializeStatsView initializes the stats view in the grid screen
func InitializeStatsView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("stats", gridX+1, 0, maxX-1, 2*minY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Game Statistics"
		v.Wrap = true
	} else {
		RefreshStatsView(v)
	}

	return nil
}

// InitializeSelectView initializes the select view in the grid screen
func InitializeSelectView(g *gocui.Gui) error {
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
		RefreshSelectView(v)
	}

	return nil
}

// InitializeSideViews initializes all side views in the grid screen
func InitializeSideViews(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if maxX < minX {
		return nil
	}
	if err := InitializeStatsView(g); err != nil {
		return err
	}
	if err := InitializeSelectView(g); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(currentView); err != nil {
		return err
	}
	return nil
}

// RefreshErrorView refreshes the error view in the grid screen
func RefreshErrorView(g *gocui.Gui, v *gocui.View) error {
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

// RefreshSelectView refreshes the select view in the grid screen
func RefreshSelectView(v *gocui.View) {
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

// RefreshSquareView refreshes a specific square on the grid screen
func RefreshSquareView(v *gocui.View) {
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

// RefreshStatsView refreshes the status view on the grid screen
func RefreshStatsView(v *gocui.View) {
	v.Clear()

	perms := 0
	fmt.Fprintf(v, "Remaining ships:\n")

	for _, ship := range h.Ships {
		fmt.Fprintf(v, "  %s\n", ship.GetType())
		perms += h.Data[ship.GetLength()].Len()
	}

	fmt.Fprintf(v, "\nTurns Taken: %d\n", h.Turns)
	fmt.Fprintf(v, "Permutations: %d\n", perms)
	fmt.Fprintf(v, "\nActive Hitstack:\n")

	for _, square := range h.HitStack {
		fmt.Fprintf(v, "%s ", square.PrintSquare())
	}

	if len(h.HitStack) == 0 {
		fmt.Fprint(v, "  Empty")
	}
}

// SwitchToGrid switches to grid view
func SwitchToGrid(g *gocui.Gui, v *gocui.View) error {
	currentView = "select"
	maxX, maxY := g.Size()

	if minX <= maxX && minY <= maxY {
		g.SetManagerFunc(GridLayout)
		SetKeyBindings(g)
		selectPos = 0
	}

	return nil
}
