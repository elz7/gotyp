package cui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/elz7/gotyp/game"
)

func mainMenuEnter(g *gocui.Gui, v *gocui.View) error {

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

func debugF10(g *gocui.Gui, v *gocui.View) error {
	return widgetSwitcher.Toggle(WidgetDebug)
}

func mainMenuArrowUp(g *gocui.Gui, v *gocui.View) error {
	cursorUp(v)
	return nil
}

func mainMenuArrowDown(g *gocui.Gui, v *gocui.View) error {
	cursorDown(v)
	return nil
}

func gameMenuArrowUp(g *gocui.Gui, v *gocui.View) error {
	c := cursorUp(v)

	if c == len(game.GameModes) {
		return setViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return setViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
}

func gameMenuArrowDown(g *gocui.Gui, v *gocui.View) error {
	c := cursorDown(v)

	if c == len(game.GameModes) {
		return setViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return setViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
}

func gameMenuEnter(g *gocui.Gui, v *gocui.View) error {
	c := cursorPos(v)

	if c == len(game.GameModes) {
		widgetSwitcher.Switch(WidgetMainMenu)
		return nil
	}

	// gameMode := game.GameModes[c]

	return widgetSwitcher.Switch(WidgetGame)
}

func changeViewVisibility(g *gocui.Gui, b bool, views ...string) {
	for _, it := range views {
		v, _ := g.View(it)
		v.Visible = b
	}
}

func setViewBufferString(g *gocui.Gui, view, text string) error {
	v, err := g.View(view)
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprint(v, text)
	return nil
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
