package main

import (
	"github.com/eaglerock1337/gobat/pkg/gobat"
)

func main() {
	screen := gobat.NewTerminal()
	gobat.Run(screen)
}
