package main

import "github.com/awesome-gocui/gocui"

func setTransitions(st *StateMachine[*gocui.Gui]) {

	st.AddTransition(Menu, Debug, func(g *gocui.Gui) error {
		g.SetViewOnBottom(MenuView)
		g.SetViewOnTop(DebugConsoleView)
		g.SetViewOnTop(DebugPromptView)

		g.SetCurrentView(DebugPromptView)
		return nil
	})

	st.AddTransition(Debug, Menu, func(g *gocui.Gui) error {
		g.SetViewOnBottom(DebugConsoleView)
		g.SetViewOnBottom(DebugPromptView)
		g.SetViewOnTop(MenuView)

		g.SetCurrentView(MenuView)

		return nil
	})

}
