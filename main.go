package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

// VARS!
var (
	widgetSwitcher *WidgetSwitcher
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	widgetSwitcher = NewWidgetSwitcher(g, WidgetMenu)
	addSwitches(widgetSwitcher)

	g.SetManagerFunc(layout)
	setKeybindings(g)

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}
}
