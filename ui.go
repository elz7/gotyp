package main

import (
	"fmt"
	"math"

	"github.com/awesome-gocui/gocui"
)

// Widgets
const (
	WidgetMainMenu       = "widget-main-menu"
	WidgetDebug          = "widget-debug"
	WidgetSelectGameMode = "widget-select-game-mode"
	WidgetSettings       = "widget-settings"
	WidgetGame           = "widget-game"
)

// Views
const (
	ViewApplication         = "view-application"
	ViewMainMenu            = "view-main-menu"
	ViewDebugConsole        = "view-debug-console"
	ViewDebugPrompt         = "view-debug-prompt"
	ViewGameModeMenu        = "view-game-mode-menu"
	ViewGameModeDescription = "view-game-mode-description"
	ViewGameInput           = "view-game-input"
	ViewGameBoard           = "view-game-board"
	ViewSettings            = "view-settings"
)

func addSwitches(vs *WidgetSwitcher) {
	// Debug section
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetDebug), func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewMainMenu)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSelectGameMode, WidgetDebug), func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewSelectGameModeMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetDebug, WidgetSelectGameMode), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewGameModeMenu)
		g.SetCurrentView(ViewGameModeMenu)
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
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSelectGameMode), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewGameModeMenu, ViewGameModeDescription)

		v, _ := g.View(ViewGameModeDescription)
		setViewBufferString(v, gameModes[0].Description)

		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetSelectGameMode, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewGameModeMenu, ViewGameModeDescription)
		changeViewVisibility(g, true, ViewMainMenu)
		return nil
	})
	vs.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSettings), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewGameModeMenu)
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

	if v, err := g.SetView(ViewDebugConsole, 1, 1, maxX-1, maxY/2-7, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Frame = true
		v.Autoscroll = true
		v.Visible = false
		v.Overwrite = true

		initDebugConsole(v)
	}

	if v, err := g.SetView(ViewDebugPrompt, 1, maxY/2-6, maxX-1, maxY/2-4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = true
		v.Visible = false
		v.Overwrite = true
		v.Highlight = true
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

	if v, err := g.SetView(ViewGameModeMenu, maxX/2-16, maxY/2-3, maxX/2+16, maxY/2+1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Select a Game Mode!"
		v.Frame = true
		v.SelFgColor = gocui.ColorGreen
		v.Highlight = true
		v.Visible = false

		for i, m := range gameModes {
			fmt.Fprintf(v, "%d. %v\n", i+1, m.Name)
		}
		fmt.Fprintf(v, "0. Back to Main Menu")
	}

	if v, err := g.SetView(ViewGameModeDescription, maxX/2-19, maxY/2+2, maxX/2+19, maxY/2+8, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Visible = false
		v.Wrap = true

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
	g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		mx, my := g.MousePosition()
		x0, y0, x1, y1, _ := g.ViewPosition(ViewDebugPrompt)

		if x0 <= mx && mx <= x1 && y0 <= my && my <= y1 {
			g.SetCurrentView(ViewDebugPrompt)
		}

		return nil
	})

	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowUp, gocui.ModNone, mainMenuCursorUp)
	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowDown, gocui.ModNone, mainMenuCursorDown)
	g.SetKeybinding(ViewMainMenu, gocui.KeyEnter, gocui.ModNone, selectViewMenuItem)
	g.SetKeybinding(ViewMainMenu, gocui.KeySpace, gocui.ModNone, selectViewMenuItem)

	g.SetKeybinding(ViewGameModeMenu, gocui.KeyArrowUp, gocui.ModNone, gameModeMenuCursorUp)
	g.SetKeybinding(ViewGameModeMenu, gocui.KeyArrowDown, gocui.ModNone, gameModeMenuCursorDown)
}

func selectViewMenuItem(g *gocui.Gui, v *gocui.View) error {

	switch cursorPos(v) {
	case 0:
		return widgetSwitcher.Switch(WidgetSelectGameMode)
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

func mainMenuCursorUp(g *gocui.Gui, v *gocui.View) error {
	cursorUp(v)
	return nil
}

func mainMenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	cursorDown(v)
	return nil
}

func gameModeMenuCursorUp(g *gocui.Gui, v *gocui.View) error {
	c := cursorUp(v)
	dv, _ := g.View(ViewGameModeDescription)
	if c == len(gameModes) {
		setViewBufferString(dv, "Main Menu")
	} else {
		setViewBufferString(dv, gameModes[c].Description)
	}
	return nil

}

func gameModeMenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	c := cursorDown(v)
	dv, _ := g.View(ViewGameModeDescription)
	if c == len(gameModes) {
		setViewBufferString(dv, "Main Menu")
	} else {
		setViewBufferString(dv, gameModes[c].Description)
	}
	return nil
}

func changeViewVisibility(g *gocui.Gui, b bool, views ...string) {
	for _, it := range views {
		v, _ := g.View(it)
		v.Visible = b
	}
}

func setViewBufferString(v *gocui.View, s string) {
	v.Clear()
	fmt.Fprint(v, s)
}

func cursorUp(v *gocui.View) int {
	_, curLine := v.Cursor()
	_, curOrig := v.Origin()

	_, viewLinesCount := v.Size()
	bufferLinesCount := v.LinesHeight()

	switch {
	case curLine == 0 && curOrig == 0: // the first line
		curOrig = viewLinesCount * (int(math.Ceil(float64(bufferLinesCount)/float64(viewLinesCount))) - 1)
		curLine = viewLinesCount - curOrig

	case curLine == 0 && curOrig != 0: // the first line somewhere in the middle
		curOrig = curOrig - viewLinesCount
		curLine = viewLinesCount - 1

	default:
		curLine = curLine - 1
	}

	v.SetOrigin(0, curOrig)
	v.SetCursor(0, curLine)

	return curOrig + curLine
}

func cursorDown(v *gocui.View) int {
	_, curLine := v.Cursor()
	_, curOrig := v.Origin()

	curLine = curLine + curOrig // current line within the buffer, not the view

	_, viewLinesCount := v.Size()
	bufferLinesCount := v.LinesHeight()

	switch {
	case curLine == bufferLinesCount-1: // the last line
		curOrig = 0
		curLine = 0
	case (curLine+1)%viewLinesCount == 0: // last line somewhere in the middle
		curOrig = curOrig + viewLinesCount
		curLine = 0
	default:
		curLine = curLine + 1
	}

	v.SetOrigin(0, curOrig)
	v.SetCursor(0, curLine)

	return curOrig + curLine
}

func cursorPos(v *gocui.View) int {
	_, y := v.Cursor()
	_, o := v.Origin()
	return y + o
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
