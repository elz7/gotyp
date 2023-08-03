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
	dv, _ := g.View(ViewGameModeDescription)
	if c == len(game.GameModes) {
		setViewBufferString(dv, "Main Menu")
	} else {
		setViewBufferString(dv, game.GameModes[c].Description)
	}

	return nil
}

func gameMenuArrowDown(g *gocui.Gui, v *gocui.View) error {
	c := cursorDown(v)
	dv, _ := g.View(ViewGameModeDescription)
	if c == len(game.GameModes) {
		setViewBufferString(dv, "Main Menu")
	} else {
		setViewBufferString(dv, game.GameModes[c].Description)
	}
	return nil
}

func gameMenuEnter(g *gocui.Gui, v *gocui.View) error {
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

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
