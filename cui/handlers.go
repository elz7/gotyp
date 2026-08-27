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

func debugArrowUp(g *gocui.Gui, v *gocui.View) error {
	v, _ = g.View(ViewDebugConsole)
	v.Autoscroll = false
	x, y := v.Origin()
	v.SetOrigin(x, y-1)
	return nil
}

func debugArrowDown(g *gocui.Gui, v *gocui.View) error {
	v, _ = g.View(ViewDebugConsole)
	v.Autoscroll = false
	x, y := v.Origin()
	v.SetOrigin(x, y+1)
	return nil
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
		return SetViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return SetViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
}

func gameMenuArrowDown(g *gocui.Gui, v *gocui.View) error {
	c := cursorDown(v)

	if c >= len(game.GameModes) {
		return SetViewBufferString(g, ViewGameModeDescription, "Back to main menu.")
	}

	return SetViewBufferString(g, ViewGameModeDescription, game.GameModes[c].Description)
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
	game.CurrentGame.GenerateGameData(board)

	go timer(g)
	go func() {
		<-syncGameOver
		widgetSwitcher.Switch(WidgetGameScore)
		c, w := game.CurrentGame.Score()
		v, _ := g.View(ViewGameScore)
		fmt.Fprintf(v, "Correct = %d\nWrong = %d", c, w)
	}()

	return widgetSwitcher.Switch(WidgetGame)
}

func gameInputEnter(g *gocui.Gui, v *gocui.View) error {
	input := v.Buffer()
	defer v.Clear()
	b, _ := g.View(ViewGameBoard)
	game.CurrentGame.PlayerMove(b, input)
	return nil
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}
