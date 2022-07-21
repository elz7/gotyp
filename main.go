package main

import (
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
)

// States
const (
	Menu  = "menu"
	Debug = "debug"
	Game  = "game"
)

// Views
const (
	MenuView         = "menu-view"
	DebugConsoleView = "debug-console-view"
	DebugPromptView  = "debug-prompt-view"
	GameInputView    = "game-input"
	GameBoardView    = "game-board"
)

// VARS!
var (
	stateMachine *StateMachine[*gocui.Gui]
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	stateMachine = NewStateMachine(g, Menu)
	setTransitions(stateMachine, g)

	g.SetManagerFunc(layout)

	if err := setKeybindings(g); err != nil {
		log.Panic(err)
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 1, 1, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "GoTyp"
		v.Frame = true
	}

	if v, err := g.SetView(DebugConsoleView, 1, 1, maxX-1, maxY-4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Frame = true
		v.Autoscroll = true
		v.Visible = false
	}

	if v, err := g.SetView(DebugPromptView, 1, maxY-3, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = true
		v.Visible = false
	}

	if v, err := g.SetView(MenuView, maxX/2-11, maxY/2-2, maxX/2+11, maxY/2+2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Menu"
		v.Frame = true
		v.SelFgColor = gocui.ColorGreen
		v.Highlight = true

		fmt.Fprintf(v, "1. Play\n")
		fmt.Fprintf(v, "2. Settings\n")
		fmt.Fprintf(v, "3. Quit")

		g.SetCurrentView(MenuView)
	}

	if v, err := g.SetView(GameInputView, maxX/2-23, maxY/2-3, maxX/2+23, maxY/2-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Editable = true
		v.Visible = false
	}

	if v, err := g.SetView(GameBoardView, maxX/2-23, maxY/2, maxX/2+23, maxY/2+6, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Visible = false
	}

	return nil
}

func setKeybindings(g *gocui.Gui) error {
	var err error
	f := func(v string, k gocui.Key, m gocui.Modifier, h func(*gocui.Gui, *gocui.View) error) bool {
		err = g.SetKeybinding(v, k, m, h)
		return err == nil
	}

	if f("", gocui.KeyCtrlC, gocui.ModNone, quit) && f("", gocui.KeyF10, gocui.ModNone, toggleDebug) &&
		f(MenuView, gocui.KeyArrowUp, gocui.ModNone, cursorUp) && f(MenuView, gocui.KeyArrowDown, gocui.ModNone, cursorDown) &&
		f(MenuView, gocui.KeyEnter, gocui.ModNone, selectMenuItem) && f(MenuView, gocui.KeySpace, gocui.ModNone, selectMenuItem) {
		return nil
	}

	return err
}

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if err := v.SetCursor(0, cy-1); err != nil {
		v.SetCursor(0, len(v.BufferLines())-1)
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	cy = (cy + 1) % len(v.BufferLines())
	v.SetCursor(0, cy)
	return nil
}

func selectMenuItem(g *gocui.Gui, v *gocui.View) error {

	switch _, i := v.Cursor(); i {
	case 0:
		return nil
	case 1:
		return nil
	case 2:
		return quit(g, v)
	}

	return nil
}

func toggleDebug(g *gocui.Gui, v *gocui.View) error {
	return stateMachine.Toggle(Debug)
}
