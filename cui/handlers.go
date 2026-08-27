package cui

import (
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

	if c >= len(game.GameModes) {
		return setViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return setViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
}

func gameMenuArrowDown(g *gocui.Gui, v *gocui.View) error {
	c := cursorDown(v)

	if c >= len(game.GameModes) {
		return setViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return setViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
}

func gameMenuEnter(g *gocui.Gui, v *gocui.View) error {
	c := cursorPos(v)

	// Back to main menu
	if c >= len(game.GameModes) {
		widgetSwitcher.Switch(WidgetMainMenu)
		return nil
	}

	gameMode := game.GameModes[c]
	game.CurrentGame = gameMode.CreateGame()

	board, _ := g.View(ViewGameBoard)
	w, h := board.Size()
	data := game.CurrentGame.GenerateGameData(w-1, h)
	setViewBufferString(g, ViewGameBoard, data)

	return widgetSwitcher.Switch(WidgetGame)
}

func gameInputEnter(g *gocui.Gui, v *gocui.View) error {
	input := v.Buffer()
	defer v.Clear()
	game.CurrentGame.PlayerMove(input)
	return nil
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
