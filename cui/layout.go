package cui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/elz7/gotyp/game"
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

		for i, m := range game.GameModes {
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

	if v, err := g.SetView(ViewGameInput, maxX/2-25, maxY/2-5, maxX/2+25, maxY/2-3, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Visible = false
		v.Editable = true
	}

	if v, err := g.SetView(ViewGameBoard, maxX/2-25, maxY/2-2, maxX/2+25, maxY/2+6, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Visible = false
	}

	return nil
}
