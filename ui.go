package main

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// Widgets
const (
	WidgetMainMenu           = "widget-main-menu"
	WidgetDebug              = "widget-debug"
	WidgetSelectGameModeMenu = "widget-select-game-mode-menu"
	WidgetSettings           = "widget-settings"
	WidgetGame               = "widget-game"
)

// Views
const (
	ViewApplication        = "view-application"
	ViewMainMenu           = "view-main-menu"
	ViewDebugConsole       = "view-debug-console"
	ViewDebugPrompt        = "view-debug-prompt"
	ViewSelectGameModeMenu = "view-select-game-mode-menu"
	ViewGameDescription    = "view-game-description"
	ViewGameInput          = "view-game-input"
	ViewGameBoard          = "view-game-board"
	ViewSettings           = "view-settings"
)

func addSwitches(vs *WidgetSwitcher) {
	// Debug section
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetDebug), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewMainMenu)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSelectGameModeMenu, WidgetDebug), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSelectGameModeMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetSelectGameModeMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewSelectGameModeMenu)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSettings, WidgetDebug), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetSettings), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	// End of debug section
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSelectGameModeMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewSelectGameModeMenu)
		g.SetCurrentView(ViewSelectGameModeMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSelectGameModeMenu, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSelectGameModeMenu)
		changeViewVisibility(g, true, ViewMainMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSettings), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewSelectGameModeMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSettings, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewMainMenu)
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

	if v, err := g.SetView(ViewMainMenu, maxX/2-11, maxY/2-2, maxX/2+11, maxY/2+2, 0); err != nil {
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

		g.SetCurrentView(ViewMainMenu)
	}

	if v, err := g.SetView(ViewSelectGameModeMenu, maxX/2-11, maxY/2-2, maxX/2+11, maxY/2+2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Select Game Mode!"
		v.Frame = true
		v.SelFgColor = gocui.ColorGreen
		v.Highlight = true
		v.Visible = false
	}

	if v, err := g.SetView(ViewSettings, 1, 1, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Settings"
		v.Frame = true
		v.SelFgColor = gocui.ColorGreen
		v.Highlight = true
		v.Visible = false
	}

	return nil
}

func setKeybindings(g *gocui.Gui) {
	// global
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, toggleWidgetDebug)

	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowUp, gocui.ModNone, cursorUp)
	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowDown, gocui.ModNone, cursorDown)
	g.SetKeybinding(ViewMainMenu, gocui.KeyEnter, gocui.ModNone, selectViewMenuItem)
	g.SetKeybinding(ViewMainMenu, gocui.KeySpace, gocui.ModNone, selectViewMenuItem)
}

func selectViewMenuItem(g *gocui.Gui, v *gocui.View) error {

	switch _, i := v.Cursor(); i {
	case 0:
		return widgetSwitcher.Switch(WidgetSelectGameModeMenu)
	case 1:
		return widgetSwitcher.Switch(WidgetSettings)
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
