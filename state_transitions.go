package main

import "github.com/awesome-gocui/gocui"

func setTransitions(st *StateMachine[*gocui.Gui], g *gocui.Gui) {

	view := func(viewname string) *gocui.View {
		v, _ := g.View(viewname)
		return v
	}

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

	st.AddTransition(Transition[*gocui.Gui]{
		From: Menu,
		To:   PreGame,
		ForwardFunc: func(g *gocui.Gui) error {
			visible(false, MenuView)
			visible(true, PreGameMenuView, PreGameDescriptionView)
			g.SetCurrentView(PreGameMenuView)

			view(PreGameMenuView).SetCursor(0, 0)
			showDescription(g, 0)

			return nil
		},
		BackwardFunc: func(g *gocui.Gui) error {
			visible(false, PreGameMenuView, PreGameDescriptionView)
			visible(true, MenuView)
			g.SetCurrentView(MenuView)
			return nil
		},
	})
}
