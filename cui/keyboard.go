package cui

import "github.com/awesome-gocui/gocui"

func setKeybindings(g *gocui.Gui) {
	// global
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, debugF10)
	g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		mx, my := g.MousePosition()
		x0, y0, x1, y1, _ := g.ViewPosition(ViewDebugPrompt)

		if x0 <= mx && mx <= x1 && y0 <= my && my <= y1 {
			g.SetCurrentView(ViewDebugPrompt)
		}

		return nil
	})

	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowUp, gocui.ModNone, mainMenuArrowUp)
	g.SetKeybinding(ViewMainMenu, gocui.KeyArrowDown, gocui.ModNone, mainMenuArrowDown)
	g.SetKeybinding(ViewMainMenu, gocui.KeyEnter, gocui.ModNone, mainMenuEnter)

	g.SetKeybinding(ViewGameModeMenu, gocui.KeyArrowUp, gocui.ModNone, gameMenuArrowUp)
	g.SetKeybinding(ViewGameModeMenu, gocui.KeyArrowDown, gocui.ModNone, gameMenuArrowDown)
	g.SetKeybinding(ViewGameModeMenu, gocui.KeyEnter, gocui.ModNone, gameMenuEnter)
}
