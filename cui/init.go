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
	// Debug section
	ws.AddSwitch(NewSwitch(WidgetMainMenu, WidgetDebug), func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetDebug, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewMainMenu)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetSelectGameMode, WidgetDebug), func(g *gocui.Gui) error {
		// changeViewVisibility(g, false, ViewSelectGameModeMenu)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		// g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetDebug, WidgetSelectGameMode), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		// changeViewVisibility(g, true, ViewGameModeMenu)
		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})

	ws.AddSwitch(NewSwitch(WidgetSettings, WidgetDebug), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewDebugConsole, ViewDebugPrompt)
		g.SetCurrentView(ViewDebugPrompt)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetDebug, WidgetSettings), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewDebugConsole, ViewDebugPrompt)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewMainMenu)
		return nil
	})
	// End of debug section
	ws.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSelectGameMode), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewGameModeMenu, ViewGameModeDescription)

		v, _ := g.View(ViewGameModeDescription)
		setViewBufferString(v, game.GameModes[0].Description)

		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetSelectGameMode, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewGameModeMenu, ViewGameModeDescription)
		changeViewVisibility(g, true, ViewMainMenu)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetMainMenu, WidgetSettings), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewMainMenu)
		changeViewVisibility(g, true, ViewSettings)
		g.SetCurrentView(ViewGameModeMenu)
		return nil
	})
	ws.AddSwitch(NewSwitch(WidgetSettings, WidgetMainMenu), func(g *gocui.Gui) error {
		changeViewVisibility(g, false, ViewSettings)
		changeViewVisibility(g, true, ViewMainMenu)
		return nil
	})
}
