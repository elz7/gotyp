package cui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/elz7/gotyp/game"
)

var widgetSwitcher *WidgetSwitcher

func Initialize(g *gocui.Gui) {
	widgetSwitcher = NewWidgetSwitcher(g, WidgetMainMenu)
	addWidgetSwitches(widgetSwitcher)

	g.SetManagerFunc(layout)
	setKeybindings(g)
}

func addWidgetSwitches(ws *WidgetSwitcher) {
	/* Debug */
	ws.AddSwitch(WidgetMainMenu, WidgetDebug, func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(WidgetDebug, WidgetMainMenu, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewMainMenu)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	ws.AddSwitch(WidgetSelectGameMode, WidgetDebug, func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewSelectGameModeMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(WidgetDebug, WidgetSelectGameMode, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewGameModeMenu)
		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	ws.AddSwitch(WidgetSettings, WidgetDebug, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(WidgetDebug, WidgetSettings, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	/* End of Debug */

	ws.AddSwitch(WidgetMainMenu, WidgetSelectGameMode, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewGameModeMenu, ViewGameModeDescription)

		v, _ := g.View(ViewGameModeMenu)
		v.SetOrigin(0, 0)
		v.SetCursor(0, 0)
		setViewBufferString(g, ViewGameModeDescription, game.GameModes[0].Description)

		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	ws.AddSwitch(WidgetSelectGameMode, WidgetMainMenu, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewGameModeMenu, ViewGameModeDescription)
		changeViewVisibility(g, true, ViewMainMenu)

		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	ws.AddSwitch(WidgetMainMenu, WidgetSettings, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	ws.AddSwitch(WidgetSettings, WidgetMainMenu, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewMainMenu)
		return nil
	})
	ws.AddSwitch(WidgetSelectGameMode, WidgetGame, func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewGameModeMenu, ViewGameModeDescription)
		changeViewVisibility(g, true, ViewGameInput, ViewGameBoard)

		g.SetCurrentView(ViewGameInput)
		return nil
	})
}
