package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	if err := setKeybindings(g); err != nil {
		log.Panic(err)
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 1, 1, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "GoTyp"
		v.Frame = true

		g := color.New(color.FgGreen)
		g.Fprint(v, "Hello World")
	}
	return nil
}

func setKeybindings(g *gocui.Gui) error {
	var err error
	f := func(v string, k gocui.Key, m gocui.Modifier, h func(*gocui.Gui, *gocui.View) error) bool {
		if err = g.SetKeybinding(v, k, m, h); err != nil {
			return false
		}
		return true
	}

	if f("", gocui.KeyCtrlC, gocui.ModNone, quit) {
		return nil
	}

	return err
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
