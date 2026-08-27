package cui

import (
	"fmt"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/elz7/gotyp/game"
)

var syncGameOver chan struct{}

func formatDuration(d time.Duration) string {
	if d.Minutes() == 0 {
		return fmt.Sprintf("00:%02d", int(d.Seconds()))
	}
	return fmt.Sprintf("%02d:%02d", int(d.Minutes()), int(d.Seconds())%60)
}

func timer(g *gocui.Gui) {
	dur := game.Settings.RoundDuration
	syncGameOver = make(chan struct{})

	for range time.Tick(time.Second) {
		title := formatDuration(dur)
		g.Update(func(g *gocui.Gui) error {
			v, _ := g.View(ViewGameInput)
			v.Title = title
			return nil
		})
		dur -= time.Second
		if dur.Seconds() < 0 {
			syncGameOver <- struct{}{}
			break
		}
	}
}
