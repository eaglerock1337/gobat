package main

import (
	"log"

	"github.com/eaglerock1337/gobat/pkg/terminal"
	"github.com/jroimartin/gocui"
)

func main() {
	screen := terminal.NewTerminal()
	defer screen.Screen.Close()

	if err := screen.Screen.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
