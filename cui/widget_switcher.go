package cui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type Widget string

type Switch string

type SwitchFunc func(*gocui.Gui) error

type WidgetSwitcher struct {
	gui       *gocui.Gui
	prev, cur Widget
	switches  map[Switch]SwitchFunc
}

func NewWidgetSwitcher(g *gocui.Gui, initialWidget Widget) *WidgetSwitcher {
	sw := make(map[Switch]SwitchFunc)
	return &WidgetSwitcher{gui: g, cur: initialWidget, switches: sw}
}

func (ws *WidgetSwitcher) AddSwitch(s Switch, f SwitchFunc) {
	ws.switches[s] = f
}

func (ws *WidgetSwitcher) Switch(to Widget) error {
	s := NewSwitch(ws.cur, to)
	if f, ok := ws.switches[s]; ok {
		ws.prev = ws.cur
		ws.cur = to
		f(ws.gui)
		return nil
	}
	return fmt.Errorf("error: switch not exists (%v)", s)
}

func (ws *WidgetSwitcher) Toggle(to Widget) error {

	cur := ws.cur
	if cur == to {
		cur = ws.prev
		if cur == "" {
			return fmt.Errorf("error: you cannot toggle from %q", to)
		}
	}
	s1, s2 := NewSwitch(cur, to), NewSwitch(to, cur)
	_, ok1 := ws.switches[s1]
	_, ok2 := ws.switches[s2]
	if !ok1 || !ok2 {
		return fmt.Errorf("error: switch (%v) and (%v) must be added", s1, s2)
	}

	if ws.cur != to {
		return ws.Switch(to)
	}

	return ws.Switch(ws.prev)
}

func NewSwitch(from, to Widget) Switch {
	return Switch(fmt.Sprintf("%q->%q", from, to))
}
