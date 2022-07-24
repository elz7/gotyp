package main

import "github.com/awesome-gocui/gocui"

func setTransitions(st *StateMachine[*gocui.Gui], g *gocui.Gui) {

	visible := func(visibility bool, views ...string) {
		for _, vn := range views {
			v, _ := g.View(vn)
			v.Visible = visibility
		}
	}

	st.AddTransition(Transition[*gocui.Gui]{
		From: Menu,
		To:   Debug,
		ForwardFunc: func(g *gocui.Gui) error {
			visible(false, MenuView)
			visible(true, DebugConsoleView, DebugPromptView)
			g.SetCurrentView(DebugPromptView)
			return nil
		},
		BackwardFunc: func(g *gocui.Gui) error {
			visible(false, DebugPromptView, DebugConsoleView)
			visible(true, MenuView)
			g.SetCurrentView(MenuView)
			return nil
		},
	})
}
