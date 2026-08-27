package cui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func changeViewVisibility(g *gocui.Gui, b bool, views ...string) {
	for _, it := range views {
		v, _ := g.View(it)
		v.Visible = b
	}
}

func setViewBufferString(g *gocui.Gui, view, text string) error {
	v, err := g.View(view)
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprint(v, text)
	return nil
}
