package main

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// Widgets
const (
	WidgetMenu        = "widget-menu"
	WidgetDebug       = "widget-debug"
	WidgetPreGameMenu = "widget-pre-game-menu"
	WidgetGame        = "widget-game"
)

// Views
const (
	ViewApplication     = "view-application"
	ViewMenu            = "view-menu"
	ViewDebugConsole    = "view-debug-console"
	ViewDebugPrompt     = "view-debug-prompt"
	ViewPreGameMenu     = "view-pre-game-menu"
	ViewGameDescription = "view-game-description"
	ViewGameInput       = "view-game-input"
	ViewGameBoard       = "view-game-board"
)

func addSwitches(vs *WidgetSwitcher) {
	vs.AddSwitch(NewSwitch(WidgetMenu, WidgetDebug), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewMenu)
		g.SetCurrentView(ViewMenu)
		return nil
	})
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(ViewApplication, 1, 1, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "GoTyp"
		v.Frame = true
	}

	if v, err := g.SetView(ViewDebugConsole, 1, 1, maxX-1, maxY-4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Frame = true
		v.Autoscroll = true
		v.Visible = false
		v.Overwrite = true
	}

	if v, err := g.SetView(ViewDebugPrompt, 1, maxY-3, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = true
		v.Visible = false
		v.Overwrite = true
	}

	if v, err := g.SetView(ViewMenu, maxX/2-11, maxY/2-2, maxX/2+11, maxY/2+2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Menu"
		v.Frame = true
		v.SelFgColor = gocui.ColorGreen
		v.Highlight = true

		fmt.Fprintf(v, "1. Play\n")
		fmt.Fprintf(v, "2. Settings\n")
		fmt.Fprintf(v, "3. Quit")

		g.SetCurrentView(ViewMenu)
	}

	return nil
}

func setKeybindings(g *gocui.Gui) {
	// global
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, toggleWidgetDebug)

	g.SetKeybinding(ViewMenu, gocui.KeyArrowUp, gocui.ModNone, cursorUp)
	g.SetKeybinding(ViewMenu, gocui.KeyArrowDown, gocui.ModNone, cursorDown)
	g.SetKeybinding(ViewMenu, gocui.KeyEnter, gocui.ModNone, selectViewMenuItem)
	g.SetKeybinding(ViewMenu, gocui.KeySpace, gocui.ModNone, selectViewMenuItem)
}

func selectViewMenuItem(g *gocui.Gui, v *gocui.View) error {

	switch _, i := v.Cursor(); i {
	case 0:
		return nil
	case 1:
		return nil
	case 2:
		return quit(g, v)
	}

	return nil
}

func toggleWidgetDebug(g *gocui.Gui, v *gocui.View) error {
	return widgetSwitcher.Toggle(WidgetDebug)
}

func changeViewVisibility(g *gocui.Gui, b bool, views ...string) {
	for _, it := range views {
		v, _ := g.View(it)
		v.Visible = b
	}
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if err := v.SetCursor(0, cy-1); err != nil {
		v.SetCursor(0, len(v.BufferLines())-1)
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	cy = (cy + 1) % len(v.BufferLines())
	v.SetCursor(0, cy)
	return nil
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
