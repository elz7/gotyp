package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/elz7/gotyp/cui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	g.Mouse = true

	cui.Initialize(g)

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}
}
